import React, {useEffect, useState} from "react";
import {useParams} from "react-router-dom";
import {Button} from "flowbite-react"; // Для получения ID вакансии из URL
import {apiClient} from "../../api"; // Импортируем axios

const JobDetailPage = () => {
    const {jobId} = useParams();
    const [job, setJob] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [userId, setUserId] = useState(null);
    const [applicationStatus, setApplicationStatus] = useState(null);
    const [userData, setUserData] = useState({});
    const fetchJobDetails = async () => {
        try {
            const response = await apiClient.get(`/jobs/${jobId}`);
            setJob(response.data);
        } catch (error) {
            setError(error.response ? error.response.data : 'Ошибка при загрузке вакансии');
        } finally {
            setLoading(false);
        }
    };
    useEffect(() => {
        fetchJobDetails();
    }, []);

    const handleApply = async () => {
        try {
            const userProfile = JSON.parse(localStorage.getItem('userData')) || {};
            setUserId(userProfile.id);
            setUserData(userProfile)
            if (userId === null) {
                alert('Вы должны быть зарегистрированы для таких действий')
            }
            const studentId = userProfile.id;
            const companyId = job.EmployerId;
            const status = "pending";

            const applicationData = {
                companyId,
                status,
                studentId,
            };
            let access_token = localStorage.getItem('accessToken');
            const response = await apiClient.post("/applications", applicationData, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${access_token}`,
                },
            });
            setApplicationStatus("Заявка отправлена!");
        } catch (error) {
            setApplicationStatus("Ошибка при отправке заявки.");
        }
    };

    if (loading) {
        return <p>Загрузка...</p>;
    }

    if (error) {
        return <p>{error}</p>;
    }

    if (!job) {
        return <p>Вакансия не найдена!</p>;
    }

    return (
        <div className="min-h-screen bg-gradient-to-r from-indigo-50 to-purple-100">
            <div className="bg-gradient-to-r from-blue-600 to-purple-600 text-white py-12 shadow-md">
                <div className="container mx-auto px-6 text-center">
                    <h1 className="text-4xl font-extrabold">{job.Title}</h1>
                    <p className="mt-2 text-lg">{job.employerId}</p>
                </div>
            </div>

            <div className="container mx-auto px-6 py-12">
                <div className="bg-white p-8 rounded-lg shadow-xl space-y-6">
                    <h2 className="text-xl font-semibold text-indigo-600">Локация:</h2>
                    <h3 className="mt-3 text-lg font-medium text-gray-800">{job.Location}</h3>
                    <h2 className="text-xl font-semibold text-indigo-600">Описание</h2>
                    <p>{job.Description}</p>
                    <Button
                        className="mt-6 bg-indigo-600 hover:bg-indigo-700 transition-all duration-300"
                        onClick={handleApply} // Обработчик нажатия
                    >
                        Подать заявку
                    </Button>
                    {applicationStatus && (
                        <p className="mt-4 text-lg font-medium text-gray-700">{applicationStatus}</p>
                    )}
                </div>
            </div>
        </div>
    );
};

export default JobDetailPage;
