import React, { useState } from 'react';
import { Button, Dropdown } from 'flowbite-react';

const CoursePageForStudent = () => {
  const course = {
    title: 'Основы веб-разработки',
    syllabus: [
      {
        topic: 'Введение в веб-разработку',
        tasks: [
          {
            taskId: 1,
            title: 'Задание 1.1',
            description: 'Изучите основные концепции веб-разработки и создайте простую страницу.',
            theoryFile: 'https://example.com/theory-1.pdf',
          },
          {
            taskId: 2,
            title: 'Задание 1.2',
            description: 'Попрактикуйтесь в создании элементов интерфейса с помощью HTML и CSS.',
            theoryFile: 'https://example.com/theory-2.pdf',
          },
        ],
      },
      {
        topic: 'Основы HTML и CSS',
        tasks: [
          {
            taskId: 3,
            title: 'Задание 2.1',
            description: 'Создайте веб-страницу с использованием различных элементов HTML.',
            theoryFile: 'https://example.com/theory-3.pdf',
          },
        ],
      },
    ],
  };

  const [selectedTask, setSelectedTask] = useState(null);
  const [file, setFile] = useState(null);

  const handleTaskSelect = (task) => {
    setSelectedTask(task);
  };

  const handleFileChange = (e) => {
    setFile(e.target.files[0]);
  };

  const handleSubmit = () => {
    if (!file) {
      alert('Пожалуйста, загрузите файл с ответом.');
    } else {
      alert('Задание успешно отправлено!');
    }
  };

  return (
    <div className="flex min-h-screen bg-gray-50">
      <div className="w-1/4 bg-gradient-to-r from-indigo-500 to-purple-600 p-6 text-white shadow-lg">
        <h2 className="text-2xl font-semibold mb-4">Темы курса</h2>
        <div className="space-y-6">
          {course.syllabus.map((topic, index) => (
            <div key={index}>
              <p className="text-lg font-semibold">{topic.topic}</p>
              <Dropdown label="Задания" className="w-full max-w-xs">
                {topic.tasks.map((task) => (
                  <Dropdown.Item
                    key={task.taskId}
                    onClick={() => handleTaskSelect(task)}
                    className="cursor-pointer hover:bg-indigo-700"
                  >
                    {task.title}
                  </Dropdown.Item>
                ))}
              </Dropdown>
            </div>
          ))}
        </div>
      </div>

      <div className="flex-1 p-8 bg-white rounded-lg shadow-xl overflow-hidden">
        {selectedTask ? (
          <div>
            <h3 className="text-3xl font-semibold text-gray-800 mb-4">{selectedTask.title}</h3>
            <p className="text-lg text-gray-700 mb-6">{selectedTask.description}</p>

            <div className="bg-gray-50 p-6 rounded-lg shadow-md mb-6">
              <h4 className="text-2xl font-medium text-gray-800 mb-4">Теория</h4>
              <a
                href={selectedTask.theoryFile}
                className="text-blue-600 hover:underline"
                target="_blank"
                rel="noopener noreferrer"
              >
                Скачайте теорию к заданию
              </a>
            </div>

            <div className="space-y-4">
              <h4 className="text-2xl font-medium text-gray-800 mb-4">Отправьте свой ответ</h4>

              <div className="flex items-center space-x-4 mb-6">
                <input
                  type="file"
                  onChange={handleFileChange}
                  className="border border-gray-300 rounded-md p-3 w-full bg-gray-50 focus:outline-none focus:ring-2 focus:ring-indigo-500"
                />
              </div>

              <Button
                onClick={handleSubmit}
                className="w-full bg-indigo-600 hover:bg-indigo-700 text-white py-3 rounded-lg shadow-md"
              >
                Отправить ответ
              </Button>
            </div>
          </div>
        ) : (
          <p className="text-lg text-gray-600">Выберите задание из списка слева.</p>
        )}
      </div>
    </div>
  );
};

export default CoursePageForStudent;
