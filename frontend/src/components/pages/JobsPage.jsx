import React, { useState } from "react";
import { Button, Label, TextInput, Select, Modal } from "flowbite-react";
import { Pagination } from "flowbite-react";
import {useNavigate} from "react-router-dom";

const JobsPage = () => {
  const jobs = [
    { job_id: 1, title: "Software Engineer", company: "TechCorp", location: "New York", description: "Develop and maintain web applications." },
    { job_id: 2, title: "Frontend Developer", company: "WebSolutions", location: "San Francisco", description: "Work with React and Tailwind CSS." },
    { job_id: 3, title: "Backend Developer", company: "DevHub", location: "Chicago", description: "Build and optimize APIs." },
  ];

  const [searchQuery, setSearchQuery] = useState("");
  const [selectedCity, setSelectedCity] = useState("");
  const [selectedCompany, setSelectedCompany] = useState("");
  const [currentPage, setCurrentPage] = useState(1);
  const itemsPerPage = 5;
  const [isModalOpen, setIsModalOpen] = useState(false);
  const nav = useNavigate()

  const cities = ["New York", "San Francisco", "Chicago"];
  const companies = ["TechCorp", "WebSolutions", "DevHub"];

  // Фильтрация вакансий
  const filteredJobs = jobs.filter((job) => {
    return (
      (job.title.toLowerCase().includes(searchQuery.toLowerCase()) || job.description.toLowerCase().includes(searchQuery.toLowerCase())) &&
      (selectedCity ? job.location === selectedCity : true) &&
      (selectedCompany ? job.company === selectedCompany : true)
    );
  });

  // Пагинация
  const totalPages = Math.ceil(filteredJobs.length / itemsPerPage);
  const jobsToShow = filteredJobs.slice((currentPage - 1) * itemsPerPage, currentPage * itemsPerPage);

  return (
    <div className="min-h-screen bg-gradient-to-r from-indigo-50 to-purple-100">
      <div className="bg-gradient-to-r from-blue-600 to-purple-600 text-white py-12 shadow-md">
        <div className="container mx-auto px-6 text-center">
          <h1 className="text-4xl font-extrabold">Все вакансии</h1>
          <p className="mt-2 text-lg">Ищите вакансии по вашим интересам и профессиональному пути.</p>
        </div>
      </div>

      <div className="container mx-auto px-6 py-12">
        <div className="bg-white p-8 rounded-lg shadow-xl space-y-6">
          {/* Поиск вакансий и кнопка фильтров */}
          <div className="flex items-center space-x-4">
            <div className="w-full sm:w-2/3">
              <Label htmlFor="search" className="block text-lg font-medium mb-2">Поиск вакансий</Label>
              <TextInput
                id="search"
                type="text"
                placeholder="Поиск по вакансии"
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
          {jobsToShow.map((job) => (
            <div key={job.job_id} className="bg-white p-6 rounded-lg shadow-lg transform hover:scale-105 transition-all duration-300 ease-in-out">
              <h3 className="text-xl font-semibold text-indigo-600">{job.title}</h3>
              <p className="text-sm text-gray-500">{job.company}</p>
              <p className="mt-3 text-gray-700">{job.description}</p>
              <div className="mt-4 text-gray-400">
                <span className="block">{job.location}</span>
              </div>
              <Button className="mt-4 w-full bg-indigo-600 hover:bg-indigo-700 transition-all duration-300" href={`/jobs/${job.job_id}`}>
                Просмотреть
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
        <Modal.Header>Фильтры вакансий</Modal.Header>
        <Modal.Body>
          <div className="space-y-4">
            {/* Фильтр по городу */}
            <div>
              <Label htmlFor="city" className="block text-lg font-medium mb-2">Город</Label>
              <Select
                id="city"
                value={selectedCity}
                onChange={(e) => setSelectedCity(e.target.value)}
                className="w-full border-gray-300 focus:ring-2 focus:ring-indigo-500"
              >
                <option value="">Выберите город</option>
                {cities.map((city, index) => (
                  <option key={index} value={city}>{city}</option>
                ))}
              </Select>
            </div>

            {/* Фильтр по компании */}
            <div>
              <Label htmlFor="company" className="block text-lg font-medium mb-2">Компания</Label>
              <Select
                id="company"
                value={selectedCompany}
                onChange={(e) => setSelectedCompany(e.target.value)}
                className="w-full border-gray-300 focus:ring-2 focus:ring-indigo-500"
              >
                <option value="">Выберите компанию</option>
                {companies.map((company, index) => (
                  <option key={index} value={company}>{company}</option>
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
            setSelectedCity("");
            setSelectedCompany("");
            setIsModalOpen(false);
          }} className="bg-gray-400 hover:bg-gray-500">
            Сбросить фильтры
          </Button>
        </Modal.Footer>
      </Modal>
    </div>
  );
};

export default JobsPage;
