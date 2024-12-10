import React from 'react';
import {Button, Card} from "flowbite-react";

const Main = () => {
    return (
    <div className="bg-gray-50 py-8 px-4 sm:px-6 lg:px-8">
      {/* Hero Section */}
      <div className="text-center max-w-4xl mx-auto mb-12">
        <h1 className="text-4xl font-bold text-gray-800 sm:text-5xl mb-4">
          Добро пожаловать в EduConnect
        </h1>
        <p className="text-lg text-gray-600 mb-8">
          Инновационная платформа для студентов, преподавателей и работодателей, направленная на улучшение образовательного процесса и трудоустройство.
        </p>
        <div className="flex justify-center gap-4">
          <Button color="primary" href="/courses">Посмотреть курсы</Button>
          <Button color="secondary" href="/jobs">Вакансии</Button>
        </div>
      </div>

      {/* Features Section */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        <Card className="shadow-lg">
          <h3 className="text-xl font-semibold text-gray-800 mb-4">Курсы</h3>
          <p className="text-gray-600 mb-4">
            Пройдите разнообразные курсы от квалифицированных преподавателей, чтобы развивать свои навыки и подготовиться к будущей карьере.
          </p>
          <Button href="/courses" className="w-full">Изучить курсы</Button>
        </Card>

        <Card className="shadow-lg">
          <h3 className="text-xl font-semibold text-gray-800 mb-4">Вакансии</h3>
          <p className="text-gray-600 mb-4">
            Найдите стажировки и вакансии от работодателей, готовых предложить вам шанс на карьерный рост.
          </p>
          <Button href="/jobs" className="w-full">Просмотреть вакансии</Button>
        </Card>

        <Card className="shadow-lg">
          <h3 className="text-xl font-semibold text-gray-800 mb-4">Форум</h3>
          <p className="text-gray-600 mb-4">
            Обсуждайте идеи, задавайте вопросы и общайтесь с другими студентами и преподавателями на нашем форуме.
          </p>
          <Button href="/forum" className="w-full">Перейти на форум</Button>
        </Card>
      </div>

      {/* Call to Action */}
      <div className="mt-16 text-center">
        <h2 className="text-3xl font-semibold text-gray-800 mb-4">Присоединяйтесь к EduConnect!</h2>
        <p className="text-lg text-gray-600 mb-8">
          Зарегистрируйтесь сегодня и начните использовать все возможности для вашего образовательного и карьерного роста.
        </p>
        <Button color="primary" href="/signup">Зарегистрироваться</Button>
      </div>
    </div>
  );
};

export default Main;