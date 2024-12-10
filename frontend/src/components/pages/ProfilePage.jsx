import React from 'react';
import { Card, Badge, Button } from 'flowbite-react';
import {useNavigate} from "react-router-dom";

const userProfile = {
  name: 'Иван Иванов',
  role: 'company',
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
};

const ProfilePage = () => {
  const nav = useNavigate()
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
    <div className="min-h-screen bg-gray-50">
      <div className="bg-gradient-to-r from-indigo-500 to-purple-600 text-white py-10 shadow-lg">
        <div className="container mx-auto px-6 text-center">
          <h1 className="text-4xl font-bold">{userProfile.name}</h1>
          <p className="text-lg mt-2">
            Роль: <span className="capitalize">{userProfile.role}</span>
          </p>
          <div className="mt-6">{renderActionButton()}</div>
        </div>
      </div>
      <div className="container mx-auto px-6 py-10">{renderContent()}</div>
    </div>
  );
};

export default ProfilePage;
