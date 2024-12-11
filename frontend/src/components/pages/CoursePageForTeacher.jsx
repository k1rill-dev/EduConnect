import React, { useState } from 'react';
import { Dropdown, Table, Button } from 'flowbite-react';

const CoursePageForTeacher = () => {
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
          },
          {
            taskId: 2,
            title: 'Задание 1.2',
            description: 'Попрактикуйтесь в создании элементов интерфейса с помощью HTML и CSS.',
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
          },
        ],
      },
    ],
  };

  const [selectedTask, setSelectedTask] = useState(null);
  const [submissions, setSubmissions] = useState([
    {
      id: 1,
      taskId: 1,
      studentName: 'Иван Иванов',
      submissionFile: 'https://example.com/submission-ivan.pdf',
      grade: null,
    },
    {
      id: 2,
      taskId: 1,
      studentName: 'Мария Петрова',
      submissionFile: 'https://example.com/submission-maria.pdf',
      grade: null,
    },
    {
      id: 3,
      taskId: 2,
      studentName: 'Дмитрий Смирнов',
      submissionFile: 'https://example.com/submission-dmitry.pdf',
      grade: null,
    },
  ]);

  const handleTaskSelect = (task) => {
    setSelectedTask(task);
  };

  const handleGradeChange = (id, grade) => {
    setSubmissions((prevSubmissions) => {
      return prevSubmissions.map((submission) =>
        submission.id === id ? { ...submission, grade: grade } : submission
      );
    });
  };

  const filteredSubmissions = submissions.filter(
    (submission) => selectedTask && submission.taskId === selectedTask.taskId
  );

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

            <h4 className="text-2xl font-medium text-gray-800 mb-4">Ответы студентов</h4>
            {filteredSubmissions.length > 0 ? (
              <Table striped className="w-full">
                <Table.Head>
                  <Table.HeadCell>Имя студента</Table.HeadCell>
                  <Table.HeadCell>Ссылка на ответ</Table.HeadCell>
                  <Table.HeadCell>Оценка</Table.HeadCell>
                </Table.Head>
                <Table.Body>
                  {filteredSubmissions.map((submission) => (
                    <Table.Row key={submission.id} className="bg-white">
                      <Table.Cell>{submission.studentName}</Table.Cell>
                      <Table.Cell>
                        <a
                          href={submission.submissionFile}
                          className="text-blue-600 hover:underline"
                          target="_blank"
                          rel="noopener noreferrer"
                        >
                          Скачать ответ
                        </a>
                      </Table.Cell>
                      <Table.Cell>
                        <input
                          type="number"
                          value={submission.grade || ''}
                          onChange={(e) => handleGradeChange(submission.id, e.target.value)}
                          className="border border-gray-300 rounded-md p-2 w-20 text-center focus:outline-none focus:ring-2 focus:ring-indigo-500"
                          min="0"
                          max="100"
                        />
                      </Table.Cell>
                    </Table.Row>
                  ))}
                </Table.Body>
              </Table>
            ) : (
              <p className="text-lg text-gray-600">Нет ответов на это задание.</p>
            )}
          </div>
        ) : (
          <p className="text-lg text-gray-600">Выберите задание из списка слева.</p>
        )}
      </div>
    </div>
  );
};

export default CoursePageForTeacher;
