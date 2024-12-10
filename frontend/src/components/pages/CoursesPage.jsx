import React, { useState, useEffect } from 'react';
import { Card, Button, Pagination } from 'flowbite-react';

const CoursesPage = () => {
  const [courses, setCourses] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);

  const allCourses = [
    { id: 1, title: 'Основы веб-разработки', description: 'Изучите основы HTML, CSS и JavaScript.' },
    { id: 2, title: 'React для начинающих', description: 'Начните изучать React и создавайте интерактивные веб-приложения.' },
    { id: 3, title: 'Python для анализа данных', description: 'Изучите Python для работы с данными и машинного обучения.' },
    { id: 4, title: 'Введение в искусственный интеллект', description: 'Освойте основы AI и машинного обучения.' },
    { id: 5, title: 'Основы SQL', description: 'Изучите работу с базами данных с помощью SQL.' },
    { id: 6, title: 'Кибербезопасность', description: 'Изучите основы защиты данных и предотвращения угроз безопасности.' },
    { id: 7, title: 'Веб-дизайн и UX/UI', description: 'Создавайте современные интерфейсы и улучшайте пользовательский опыт.' },
    { id: 8, title: 'Дискретная математика', description: 'Изучите основы дискретной математики для IT-специалистов.' },
    { id: 9, title: 'Мобильная разработка', description: 'Изучите создание мобильных приложений для Android и iOS.' },
    { id: 10, title: 'Разработка на C++', description: 'Освойте основы программирования на C++.' },
    { id: 11, title: 'Машинное обучение', description: 'Изучите основы машинного обучения и применения моделей.' },
    { id: 12, title: 'Сетевые технологии', description: 'Научитесь основам сетевых технологий и администрирования.' },
    { id: 13, title: 'Анализ данных с помощью Excel', description: 'Овладейте анализом данных с помощью Microsoft Excel.' },
    { id: 14, title: 'Искусственный интеллект в бизнесе', description: 'Как применять AI для улучшения бизнес-процессов.' },
    { id: 15, title: 'Разработка игр', description: 'Освойте основы создания игр с использованием Unity.' }
  ];

  const coursesPerPage = 6;

  const getCoursesForCurrentPage = () => {
    const startIndex = (currentPage - 1) * coursesPerPage;
    const endIndex = startIndex + coursesPerPage;
    return allCourses.slice(startIndex, endIndex);
  };

  useEffect(() => {
    setCourses(getCoursesForCurrentPage());
    setTotalPages(Math.ceil(allCourses.length / coursesPerPage));
  }, [currentPage]);

  return (
    <div className="py-8 px-4 sm:px-6 lg:px-8 bg-gray-50">
      <h1 className="text-3xl font-bold text-center text-gray-800 mb-12">Все курсы</h1>

      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
        {courses.map((course) => (
          <Card key={course.id} className="shadow-lg">
            <h3 className="text-xl font-semibold text-gray-800 mb-4">{course.title}</h3>
            <p className="text-gray-600 mb-4">{course.description}</p>
            <Button href={`/courses/${course.id}`} className="w-full">
              Подробнее
            </Button>
          </Card>
        ))}
      </div>

      <div className="flex justify-center mt-8">
        <Pagination
          currentPage={currentPage}
          totalPages={totalPages}
          onPageChange={(page) => setCurrentPage(page)}
          className="bg-white"
        />
      </div>
    </div>
  );
};

export default CoursesPage;
