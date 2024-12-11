import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { Button } from 'flowbite-react'; // Для получения ID курса из URL

const CoursePage = () => {
  const [course, setCourse] = useState(null);
  const { courseId } = useParams(); // Получаем ID курса из параметров маршрута

  const allCourses = [
    {
      course_id: 1,
      title: 'Основы веб-разработки',
      description: 'Изучите основы HTML, CSS и JavaScript.',
      photo: 'https://via.placeholder.com/800x400?text=Course+Image',
      teacher_id: 2,
      start_date: '2024-01-15',
      end_date: '2024-04-15',
      created_at: '2023-11-20',
      syllabus: [
        'Введение в веб-разработку',
        'Основы HTML и структуры страниц',
        'Основы CSS и стилизация элементов',
        'Основы JavaScript и динамическое поведение',
        'Проект: создание веб-сайта',
      ],
    },
    {
      course_id: 2,
      title: 'React для начинающих',
      description: 'Начните изучать React и создавайте интерактивные веб-приложения.',
      photo: 'https://via.placeholder.com/800x400?text=React+Course+Image',
      teacher_id: 3,
      start_date: '2024-02-01',
      end_date: '2024-05-01',
      created_at: '2023-12-10',
      syllabus: [
        'Введение в React',
        'Компоненты и JSX',
        'Управление состоянием с помощью useState',
        'Маршрутизация с React Router',
        'Проект: создание интерактивного приложения',
      ],
    },
  ];

  useEffect(() => {
    const courseData = allCourses.find((course) => course.course_id === parseInt(courseId));
    setCourse(courseData);
  }, [courseId]);

  if (!course) {
    return (
      <div className="flex justify-center items-center py-12">
        <p className="text-xl text-gray-600">Загрузка...</p>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gradient-to-r from-indigo-50 to-purple-100">
      <div className="bg-gradient-to-r from-blue-600 to-purple-600 text-white py-12 shadow-md">
        <div className="container mx-auto px-6 text-center">
          <h1 className="text-4xl font-extrabold">{course.title}</h1>
          <p className="mt-2 text-lg">{course.description}</p>
        </div>
      </div>

      <div className="container mx-auto px-6 py-12 flex">
        {/* Боковая панель */}
        <div className="w-1/4 space-y-4">
          <div className="bg-white p-6 rounded-lg shadow-lg">
            <h2 className="text-xl font-semibold text-indigo-600">Информация о курсе</h2>
            <p><strong>Преподаватель:</strong> Преподаватель #{course.teacher_id}</p>
            <p><strong>Дата начала:</strong> {new Date(course.start_date).toLocaleDateString()}</p>
            <p><strong>Дата окончания:</strong> {new Date(course.end_date).toLocaleDateString()}</p>
            <p><strong>Дата создания:</strong> {new Date(course.created_at).toLocaleDateString()}</p>
          </div>
        </div>

        {/* Основное содержимое курса */}
        <div className="w-3/4 ml-8">
          <div className="bg-white p-8 rounded-lg shadow-xl space-y-6">
            {/* Картинка курса */}
            <div className="mb-6">
              <img className="w-full h-96 object-cover rounded-lg" src={course.photo} alt={course.title} />
            </div>
            {/* Программа курса */}
            <h2 className="text-3xl font-semibold text-gray-800 mb-4">Программа курса</h2>
            <ul className="list-inside list-disc space-y-2 text-lg text-gray-700">
              {course.syllabus.map((topic, index) => (
                <li key={index}>{topic}</li>
              ))}
            </ul>

            {/* Кнопки */}
            <div className="flex space-x-4 mt-6">
              <Button
                href={`/courses/${course.course_id}/enroll`}
                className="w-full bg-indigo-600 hover:bg-indigo-700 text-white py-3"
              >
                Записаться на курс
              </Button>
              <Button
                href={`/courseStudent/${course.course_id}/`}
                className="w-full bg-indigo-600 hover:bg-indigo-700 text-white py-3"
              >
                Перейти на курс
              </Button>
               <Button
                href={`/courseTeacher/${course.course_id}/`}
                className="w-full bg-indigo-600 hover:bg-indigo-700 text-white py-3"
              >
                Перейти на курс(преподавателю)
              </Button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default CoursePage;
