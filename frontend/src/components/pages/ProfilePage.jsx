import React, {useEffect, useState} from 'react';
import {Button, Card, Modal, Textarea, TextInput} from 'flowbite-react';
import {useNavigate} from "react-router-dom";
import {apiClient} from "../../api";

const ProfilePage = () => {
    const nav = useNavigate();

    const userProfile = JSON.parse(localStorage.getItem('userData')) || {};
    const userId = userProfile.id;

    const [isModalOpen, setIsModalOpen] = useState(false);
    const [portfolioItem, setPortfolioItem] = useState({title: '', description: '', url: ''});
    const [isEditing, setIsEditing] = useState(false);
    const [portfolio, setPortfolio] = useState([]);
    const [courses, setCourses] = useState([]);
    const [vacancies, setVacancies] = useState([]);
    const [applications, setApplications] = useState([]);

    const getPortfolio = async (studentId) => {
        try {
            const response = await apiClient.get(`/portfolios/student/${studentId}`);
            console.log(response.data.Items)
            setPortfolio(response.data.Items || []);
        } catch (error) {
            console.error("Ошибка при получении портфолио", error);
        }
    };

    const getAllCourses = async (userId) => {
        try {
            const response = await apiClient.get(`/api/courses/user/${userId}`);
            console.log(response)
            setCourses(response.data);
        } catch (error) {
            console.error("Ошибка при получении курсов пользователя", error);
        }
    };

    const addPortfolioItem = async (portfolioItem) => {
        try {
            let portfolio = [];

            const response = await apiClient.get(`/portfolios/student/${userId}`);
            portfolio = response.data.Items || [];
            if (portfolio.length > 0) {
                console.log([portfolioItem])
                const addItemResponse = await apiClient.post(
                    `/portfolios/${userId}/items`,
                    [portfolioItem],
                    {
                        headers: {
                            Authorization: `Bearer ${userProfile.access_token}`,
                        },
                    }
                );
                console.log(addItemResponse)
                setPortfolio([...portfolio, [portfolioItem]]);
                window.location.reload()
            }
        } catch (error) {
            if (error.response && error.response.status === 404) {
                const createPortfolioResponse = await apiClient.post(
                    `/portfolios`,
                    {
                        student_id: userId,
                        items: [portfolioItem],
                    },
                    {
                        headers: {
                            Authorization: `Bearer ${userProfile.access_token}`,
                        },
                    }
                );

                setPortfolio(createPortfolioResponse.data.Items);
                setIsModalOpen(false);
                setPortfolioItem({title: '', description: '', url: ''});
            } else {
                console.error("Ошибка при добавлении проекта в портфолио", error);
            }
        }
    };

    const updatePortfolioItem = async (portfolioItemId, portfolioItem) => {
        try {
            const response = await apiClient.put(`/portfolios/items/${portfolioItemId}`, portfolioItem, {
                headers: {
                    Authorization: `Bearer ${userProfile.access_token}`,
                },
            });
            setPortfolio(prevPortfolio =>
                prevPortfolio.map(item => item.id === portfolioItemId ? response.data : item)
            );
            setIsModalOpen(false);
            setPortfolioItem({title: '', description: '', url: ''});
        } catch (error) {
            console.error("Ошибка при обновлении проекта в портфолио", error);
        }
    };
    const handleAcceptResponse = async (responseId) => {
        try {
            await apiClient.put(`/applications/${responseId}/status?status=accepted`, {}, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${userProfile.access_token}`,
                },
            });
            setVacancies((prevVacancies) =>
                prevVacancies.map((vacancy) => ({
                    ...vacancy,
                    responses: vacancy.responses.map((response) =>
                        response.id === responseId ? {...response, status: 'accepted'} : response
                    ),
                }))
            );
            alert('Отклик принят');
        } catch (error) {
            console.error("Ошибка при принятии отклика", error);
        }
    };

    const handleRejectResponse = async (responseId) => {
        try {
            await apiClient.put(`/applications/${responseId}/status?status=rejected`);
            setVacancies((prevVacancies) =>
                prevVacancies.map((vacancy) => ({
                    ...vacancy,
                    responses: vacancy.responses.map((response) =>
                        response.id === responseId ? {...response, status: 'rejected'} : response
                    ),
                }))
            );
            alert('Отклик отклонен');
        } catch (error) {
            console.error("Ошибка при отклонении отклика", error);
        }
    };

    const getVacancies = async () => {
        try {
            const response = await apiClient.get('/jobs/search');
            setVacancies(response.data);
        } catch (error) {
            console.error("Ошибка при получении вакансий", error);
        }
    };

    const getAuthoredCourses = async () => {
        try {
            const response = await apiClient.get(`/courses/teacher/${userId}`);
            setCourses(response.data);
        } catch (error) {
            console.error("Ошибка при получении курсов учителя", error);
        }
    };

    const handlePortfolioChange = (e) => {
        const {name, value} = e.target;
        setPortfolioItem({...portfolioItem, [name]: value});
    };

    const handleSavePortfolio = () => {
        if (isEditing) {
            updatePortfolioItem(portfolioItem.id, portfolioItem);
        } else {
            addPortfolioItem(portfolioItem);
        }
    };
    const getApplications = async () => {
        try {
            let endpoint = '';
            if (userProfile.role === 'student') {
                endpoint = '/applications/student'; // Отклики, связанные с текущим студентом
            } else if (userProfile.role === 'company') {
                endpoint = '/applications/company'; // Отклики на вакансии компании
            }

            if (endpoint) {
                const response = await apiClient.get(endpoint, {
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${userProfile.access_token}`,
                    },
                });
                setApplications(response.data || []);
            } else {
                console.error("Неизвестная роль пользователя, отклики не будут загружены.");
            }
        } catch (error) {
            console.error("Ошибка при получении откликов", error);
        }
    };
    useEffect(() => {
        if (userProfile.role === 'student' || userProfile.role === 'company') {
            getApplications();
        }
        if (userProfile.role === 'student') {
            getPortfolio(userId);
            getAllCourses(userId);
        } else if (userProfile.role === 'company') {
            getVacancies();
        } else if (userProfile.role === 'teacher') {
            getAuthoredCourses();
        }
    }, [userProfile.role, userId]);
    const handleEditPortfolioItem = (item) => {
        setPortfolioItem(item);
        setIsEditing(true);
        setIsModalOpen(true);
    };

    const renderContent = () => {
        switch (userProfile.role) {
            case 'student':
                return (
                    <div>
                        <h2 className="text-2xl font-semibold mb-4">Мои курсы</h2>
                        {courses.length > 0 ? (
                            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                                {courses.map((course) => (
                                    <Card key={course.id} className="hover:shadow-lg">
                                        <h3 className="text-xl font-medium text-gray-800">{course.title}</h3>
                                        <p className="text-gray-600">{course.description}</p>
                                    </Card>
                                ))}
                            </div>
                        ) : (
                            <p>Нет курсов</p>
                        )}
                        <h2 className="text-2xl font-semibold mb-4">Мои отклики</h2>
                        {applications.length > 0 ? (
                            <div className="space-y-4">
                                {applications.map((application, index) => (
                                    <Card key={index} className="hover:shadow-lg">
                                        <h3 className="text-xl font-medium text-gray-800">{application.jobTitle}</h3>
                                        <p className="text-gray-600">Компания: Вебпрактик</p>
                                        <p className="text-gray-600">Статус:
                                            <span className={`font-semibold ${
                                                application.status === 'accepted'
                                                    ? 'text-green-600'
                                                    : application.status === 'rejected'
                                                        ? 'text-red-600'
                                                        : 'text-gray-600'
                                            }`}>
                                    {application.status === 'accepted' ? 'Принято' :
                                        application.status === 'rejected' ? 'Отклонено' : 'На рассмотрении'}
                                </span>
                                        </p>
                                    </Card>
                                ))}
                            </div>
                        ) : (
                            <p>Вы еще не откликались на вакансии.</p>
                        )}
                        <h2 className="text-2xl font-semibold mb-4 mt-8">Моё портфолио</h2>
                        <div className="space-y-4">
                            {portfolio ? (
                                portfolio.map((item, index) => (
                                    <Card key={index} className="hover:shadow-lg">
                                        <h3 className="text-xl font-medium text-gray-800">{item.Title}</h3>
                                        <p className="text-gray-600">{item.Description}</p>
                                        <a href={item.URL} target="_blank" rel="noopener noreferrer"
                                           className="text-indigo-600">Перейти к проекту</a>
                                        <Button onClick={() => handleEditPortfolioItem(item)}
                                                className="mt-2 bg-indigo-600 hover:bg-indigo-700">Редактировать</Button>
                                    </Card>
                                ))
                            ) : (
                                <p>Нет элементов в портфолио.</p>
                            )}
                        </div>
                        <Button onClick={() => setIsModalOpen(true)}
                                className="mt-4 bg-indigo-600 hover:bg-indigo-700 text-white shadow-md">Добавить
                            проект</Button>
                    </div>
                );

            case 'company':
                return (
                    <div>
                        <h2 className="text-2xl font-semibold mb-4">Отклики на мои вакансии</h2>
                        {applications.length > 0 ? (
                            <div className="space-y-4">
                                {applications.map((application, index) => (
                                    <Card key={index} className="p-4 border">
                                        <h3 className="text-xl font-medium text-gray-800">
                                            {application.applicantName} ({application.applicantEmail})
                                        </h3>
                                        <p className="text-gray-600">{application.message}</p>
                                        <div className="mt-2">
                                            {application.status === 'accepted' ? (
                                                <span className="text-green-600 font-semibold">Принято</span>
                                            ) : application.status === 'rejected' ? (
                                                <span className="text-red-600 font-semibold">Отклонено</span>
                                            ) : (
                                                <div className="flex space-x-2">
                                                    <Button
                                                        className="bg-green-600 hover:bg-green-700"
                                                        onClick={() => handleAcceptResponse(application.id)}
                                                    >
                                                        Принять
                                                    </Button>
                                                    <Button
                                                        className="bg-red-600 hover:bg-red-700"
                                                        onClick={() => handleRejectResponse(application.id)}
                                                    >
                                                        Отклонить
                                                    </Button>
                                                </div>
                                            )}
                                        </div>
                                    </Card>
                                ))}
                            </div>
                        ) : (
                            <p>На ваши вакансии пока нет откликов.</p>
                        )}
                    </div>
                );
            case 'teacher':
                return (
                    <div>
                        <h2 className="text-2xl font-semibold mb-4">Курсы за авторством</h2>
                        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                            {courses.map((course) => (
                                <Card key={course.id} className="hover:shadow-lg">
                                    <h3 className="text-xl font-medium text-gray-800">{course.title}</h3>
                                    <p className="text-gray-600">{course.description}</p>
                                </Card>
                            ))}
                        </div>
                    </div>
                );
            default:
                return <p className="text-gray-600">Нет доступных данных для отображения.</p>;
        }
    };

    const renderActionButton = () => {
        if (userProfile.role === 'teacher') {
            return (
                <Button
                    onClick={() => nav('/create-course')}
                    className="bg-indigo-600 hover:bg-indigo-700 text-white shadow-md"
                >
                    Создать курс
                </Button>
            );
        }

        if (userProfile.role === 'company') {
            return (
                <Button
                    onClick={() => nav('/create-job')}
                    className="bg-purple-600 hover:bg-purple-700 text-white shadow-md"
                >
                    Создать вакансию
                </Button>
            );
        }

        return null;
    };

    return (
        <div className="min-h-screen bg-gradient-to-r from-indigo-50 to-purple-100">
            <div className="bg-gradient-to-r from-blue-600 to-purple-600 text-white py-12 shadow-md">
                <div className="container mx-auto px-6 text-center">
                    <img
                        src={userProfile.picture || 'https://via.placeholder.com/150'}
                        alt="Profile"
                        className="w-32 h-32 mx-auto rounded-full"
                    />
                    <h1 className="text-4xl font-extrabold">{userProfile.first_name} {userProfile.surname}</h1>
                    <p className="text-lg mt-2">
                        Роль: <span className="capitalize">{userProfile.role}</span>
                    </p>
                    <p className="mt-2">{userProfile.bio}</p>
                    <div className="mt-6">{renderActionButton()}</div>
                </div>
            </div>

            <div className="container mx-auto px-6 py-12">
                <div className="space-y-10">
                    {/* Блок с курсами или вакансиями в зависимости от роли */}
                    {renderContent()}
                </div>
            </div>

            {/* Модальное окно для добавления/редактирования портфолио */}
            <Modal show={isModalOpen} onClose={() => setIsModalOpen(false)}>
                <Modal.Header>{isEditing ? 'Редактировать проект' : 'Добавить проект'}</Modal.Header>
                <Modal.Body>
                    <div className="space-y-4">
                        <TextInput
                            label="Название проекта"
                            name="title"
                            value={portfolioItem.title}
                            onChange={handlePortfolioChange}
                        />
                        <Textarea
                            label="Описание"
                            name="description"
                            value={portfolioItem.description}
                            onChange={handlePortfolioChange}
                        />
                        <TextInput
                            label="Ссылка на проект"
                            name="url"
                            value={portfolioItem.url}
                            onChange={handlePortfolioChange}
                        />
                    </div>
                </Modal.Body>
                <Modal.Footer>
                    <Button onClick={handleSavePortfolio}>{isEditing ? 'Сохранить' : 'Добавить'}</Button>
                    <Button onClick={() => setIsModalOpen(false)}
                            className="bg-gray-500 hover:bg-gray-600">Отмена</Button>
                </Modal.Footer>
            </Modal>
        </div>
    );
};

export default ProfilePage;
