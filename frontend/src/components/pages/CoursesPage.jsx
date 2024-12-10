import React, { useState } from "react";
import { Button, Label, TextInput, Select, Modal } from "flowbite-react";
import { Pagination } from "flowbite-react";
import { useNavigate } from "react-router-dom";

const CoursesPage = () => {
  const courses = [
    { id: 1, title: "Основы веб-разработки", description: "Изучите основы HTML, CSS и JavaScript.", category: "Веб-разработка", level: "Начальный" },
    { id: 2, title: "React для начинающих", description: "Начните изучать React и создавайте интерактивные веб-приложения.", category: "Веб-разработка", level: "Средний" },
    { id: 3, title: "Python для анализа данных", description: "Изучите Python для работы с данными и машинного обучения.", category: "Программирование", level: "Средний" },
    { id: 4, title: "Введение в искусственный интеллект", description: "Освойте основы AI и машинного обучения.", category: "Искусственный интеллект", level: "Начальный" },
    { id: 5, title: "Основы SQL", description: "Изучите работу с базами данных с помощью SQL.", category: "Базы данных", level: "Начальный" },
    { id: 6, title: "Кибербезопасность", description: "Изучите основы защиты данных и предотвращения угроз безопасности.", category: "Безопасность", level: "Средний" },
    { id: 7, title: "Веб-дизайн и UX/UI", description: "Создавайте современные интерфейсы и улучшайте пользовательский опыт.", category: "Дизайн", level: "Средний" },
    { id: 8, title: "Дискретная математика", description: "Изучите основы дискретной математики для IT-специалистов.", category: "Математика", level: "Начальный" },
    { id: 9, title: "Мобильная разработка", description: "Изучите создание мобильных приложений для Android и iOS.", category: "Мобильная разработка", level: "Средний" },
    { id: 10, title: "Разработка на C++", description: "Освойте основы программирования на C++.", category: "Программирование", level: "Начальный" },
    { id: 11, title: "Машинное обучение", description: "Изучите основы машинного обучения и применения моделей.", category: "Искусственный интеллект", level: "Продвинутый" },
    { id: 12, title: "Сетевые технологии", description: "Научитесь основам сетевых технологий и администрирования.", category: "Сети", level: "Средний" },
    { id: 13, title: "Анализ данных с помощью Excel", description: "Овладейте анализом данных с помощью Microsoft Excel.", category: "Данные", level: "Начальный" },
    { id: 14, title: "Искусственный интеллект в бизнесе", description: "Как применять AI для улучшения бизнес-процессов.", category: "Искусственный интеллект", level: "Продвинутый" },
    { id: 15, title: "Разработка игр", description: "Освойте основы создания игр с использованием Unity.", category: "Игры", level: "Средний" }
  ];

  const [searchQuery, setSearchQuery] = useState("");
  const [selectedCategory, setSelectedCategory] = useState("");
  const [selectedLevel, setSelectedLevel] = useState("");
  const [currentPage, setCurrentPage] = useState(1);
  const itemsPerPage = 6;
  const [isModalOpen, setIsModalOpen] = useState(false);
  const navigate = useNavigate();

  const categories = ["Веб-разработка", "Программирование", "Искусственный интеллект", "Базы данных", "Безопасность", "Дизайн", "Математика", "Игры"];
  const levels = ["Начальный", "Средний", "Продвинутый"];

  // Фильтрация курсов
  const filteredCourses = courses.filter((course) => {
    return (
      (course.title.toLowerCase().includes(searchQuery.toLowerCase()) || course.description.toLowerCase().includes(searchQuery.toLowerCase())) &&
      (selectedCategory ? course.category === selectedCategory : true) &&
      (selectedLevel ? course.level === selectedLevel : true)
    );
  });

  // Пагинация
  const totalPages = Math.ceil(filteredCourses.length / itemsPerPage);
  const coursesToShow = filteredCourses.slice((currentPage - 1) * itemsPerPage, currentPage * itemsPerPage);

  return (
    <div className="min-h-screen bg-gradient-to-r from-indigo-50 to-purple-100">
      <div className="bg-gradient-to-r from-blue-600 to-purple-600 text-white py-12 shadow-md">
        <div className="container mx-auto px-6 text-center">
          <h1 className="text-4xl font-extrabold">Все курсы</h1>
          <p className="mt-2 text-lg">Ищите курсы по своим интересам и навыкам.</p>
        </div>
      </div>

      <div className="container mx-auto px-6 py-12">
        <div className="bg-white p-8 rounded-lg shadow-xl space-y-6">
          {/* Поиск курсов и кнопка фильтров */}
          <div className="flex items-center space-x-4">
            <div className="w-full sm:w-2/3">
              <Label htmlFor="search" className="block text-lg font-medium mb-2">Поиск курсов</Label>
              <TextInput
                id="search"
                type="text"
                placeholder="Поиск по курсу"
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                className="w-full border-gray-300 focus:ring-2 focus:ring-indigo-500"
              />
            </div>
            <div>
              <Button onClick={() => setIsModalOpen(true)} className="bg-indigo-600 mt-9 hover:bg-indigo-700 transition-all duration-300">
                Фильтры
              </Button>
            </div>
          </div>
        </div>

        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8 mt-8">
          {coursesToShow.map((course) => (
            <div key={course.id} className="bg-white p-6 rounded-lg shadow-lg transform hover:scale-105 transition-all duration-300 ease-in-out">
              <h3 className="text-xl font-semibold text-indigo-600">{course.title}</h3>
              <p className="mt-3 text-gray-700">{course.description}</p>
              <div className="mt-4 text-gray-400">
                <span className="block">{course.category}</span>
                <span className="block">{course.level}</span>
              </div>
              <Button className="mt-4 w-full bg-indigo-600 hover:bg-indigo-700 transition-all duration-300" onClick={() => navigate(`/courses/${course.id}`)}>
                Подробнее
              </Button>
            </div>
          ))}
        </div>

        {/* Пагинация */}
        <div className="flex justify-center mt-8">
          <Pagination
            currentPage={currentPage}
            totalPages={totalPages}
            onPageChange={(page) => setCurrentPage(page)}
            showIcons={true}
            className="space-x-4"
          />
        </div>
      </div>

      {/* Модальное окно для фильтров */}
      <Modal show={isModalOpen} onClose={() => setIsModalOpen(false)}>
        <Modal.Header>Фильтры курсов</Modal.Header>
        <Modal.Body>
          <div className="space-y-4">
            {/* Фильтр по категории */}
            <div>
              <Label htmlFor="category" className="block text-lg font-medium mb-2">Категория</Label>
              <Select
                id="category"
                value={selectedCategory}
                onChange={(e) => setSelectedCategory(e.target.value)}
                className="w-full border-gray-300 focus:ring-2 focus:ring-indigo-500"
              >
                <option value="">Выберите категорию</option>
                {categories.map((category, index) => (
                  <option key={index} value={category}>{category}</option>
                ))}
              </Select>
            </div>

            {/* Фильтр по уровню */}
            <div>
              <Label htmlFor="level" className="block text-lg font-medium mb-2">Уровень</Label>
              <Select
                id="level"
                value={selectedLevel}
                onChange={(e) => setSelectedLevel(e.target.value)}
                className="w-full border-gray-300 focus:ring-2 focus:ring-indigo-500"
              >
                <option value="">Выберите уровень</option>
                {levels.map((level, index) => (
                  <option key={index} value={level}>{level}</option>
                ))}
              </Select>
            </div>
          </div>
        </Modal.Body>
        <Modal.Footer>
          <Button onClick={() => setIsModalOpen(false)} className="bg-indigo-600 hover:bg-indigo-700">
            Применить
          </Button>
          <Button onClick={() => {
            setSelectedCategory("");
            setSelectedLevel("");
            setIsModalOpen(false);
          }} className="bg-gray-400 hover:bg-gray-500">
            Сбросить фильтры
          </Button>
        </Modal.Footer>
      </Modal>
    </div>
  );
};

export default CoursesPage;
