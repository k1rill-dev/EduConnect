import React, { useState } from 'react';
import { Card, Badge, Button, Modal, TextInput, Textarea } from 'flowbite-react';
import { useNavigate } from "react-router-dom";

const userProfile = {
  name: 'Иван Иванов',
  role: 'student',
  courses: [
    { id: 1, title: 'Основы веб-разработки', description: 'HTML, CSS, JavaScript' },
    { id: 2, title: 'Реактивное программирование', description: 'React.js, Redux' },
  ],
  vacancies: [
    { id: 1, title: 'Frontend-разработчик', company: 'TechCorp', location: 'Москва' },
    { id: 2, title: 'Backend-разработчик', company: 'Innovatech', location: 'Санкт-Петербург' },
  ],
  authoredCourses: [
    { id: 1, title: 'Основы Python', description: 'Основы языка программирования Python' },
    { id: 2, title: 'Машинное обучение', description: 'Введение в машинное обучение' },
  ],
  portfolio: {
    id: 'student123',
    items: [
      { description: 'Проект по созданию сайта', title: 'Сайт-портфолио', url: 'https://myportfolio.com' }
    ]
  }
};

const ProfilePage = () => {
  const nav = useNavigate();
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [portfolioItem, setPortfolioItem] = useState({ title: '', description: '', url: '' });
  const [isEditing, setIsEditing] = useState(false);
  const [portfolio, setPortfolio] = useState(userProfile.portfolio.items);

  const handlePortfolioChange = (e) => {
    const { name, value } = e.target;
    setPortfolioItem({ ...portfolioItem, [name]: value });
  };

  const handleSavePortfolio = () => {
    if (isEditing) {
      const updatedPortfolio = portfolio.map(item =>
        item.title === portfolioItem.title ? portfolioItem : item
      );
      setPortfolio(updatedPortfolio);
    } else {
      setPortfolio([...portfolio, portfolioItem]);
    }
    setIsModalOpen(false);
    setPortfolioItem({ title: '', description: '', url: '' });
    setIsEditing(false);
  };

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
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              {userProfile.courses.map((course) => (
                <Card key={course.id} className="hover:shadow-lg">
                  <h3 className="text-xl font-medium text-gray-800">{course.title}</h3>
                  <p className="text-gray-600">{course.description}</p>
                </Card>
              ))}
            </div>

            <h2 className="text-2xl font-semibold mb-4 mt-8">Моё портфолио</h2>
            <div className="space-y-4">
              {portfolio.length > 0 ? (
                portfolio.map((item, index) => (
                  <Card key={index} className="hover:shadow-lg">
                    <h3 className="text-xl font-medium text-gray-800">{item.title}</h3>
                    <p className="text-gray-600">{item.description}</p>
                    <a href={item.url} target="_blank" rel="noopener noreferrer" className="text-indigo-600">Перейти к проекту</a>
                    <Button onClick={() => handleEditPortfolioItem(item)} className="mt-2 bg-indigo-600 hover:bg-indigo-700">Редактировать</Button>
                  </Card>
                ))
              ) : (
                <p>Нет элементов в портфолио.</p>
              )}
            </div>
            <Button onClick={() => setIsModalOpen(true)} className="mt-4 bg-indigo-600 hover:bg-indigo-700 text-white shadow-md">Добавить проект</Button>
          </div>
        );
      case 'company':
        return (
          <div>
            <h2 className="text-2xl font-semibold mb-4">Мои вакансии</h2>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              {userProfile.vacancies.map((vacancy) => (
                <Card key={vacancy.id} className="hover:shadow-lg">
                  <h3 className="text-xl font-medium text-gray-800">{vacancy.title}</h3>
                  <p className="text-gray-600">{vacancy.company}</p>
                  <Badge color="info">{vacancy.location}</Badge>
                </Card>
              ))}
            </div>
          </div>
        );
      case 'teacher':
        return (
          <div>
            <h2 className="text-2xl font-semibold mb-4">Курсы за авторством</h2>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              {userProfile.authoredCourses.map((course) => (
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
          <h1 className="text-4xl font-extrabold">{userProfile.name}</h1>
          <p className="text-lg mt-2">
            Роль: <span className="capitalize">{userProfile.role}</span>
          </p>
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
          <Button onClick={() => setIsModalOpen(false)} className="bg-gray-500 hover:bg-gray-600">Отмена</Button>
        </Modal.Footer>
      </Modal>
    </div>
  );
};

export default ProfilePage;
