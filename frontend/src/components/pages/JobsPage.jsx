import React, {useEffect, useState} from "react";
import {Button, Label, Modal, Pagination, Select, TextInput} from "flowbite-react";
import {useNavigate} from "react-router-dom";
import {apiClient} from "../../api";

const JobsPage = () => {
    const [jobs, setJobs] = useState([]);
    const [searchQuery, setSearchQuery] = useState("");
    const [selectedCity, setSelectedCity] = useState("");
    const [selectedCompany, setSelectedCompany] = useState("");
    const [currentPage, setCurrentPage] = useState(1);
    const [totalPages, setTotalPages] = useState(1);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const nav = useNavigate();

    const itemsPerPage = 40;
    const cities = ["New York", "San Francisco", "Chicago"];
    const companies = ["TechCorp", "WebSolutions", "DevHub"];

    const fetchJobs = async (page) => {
    try {
        const response = await apiClient.get(`/jobs/filter`, {
            params: {
                page,
                limit: itemsPerPage,
            }
        });
        console.log(response.data)
        const jobsData = response.data || [];
        setJobs(jobsData);

        const totalJobs = jobsData.length;
        const totalPages = Math.max(1, Math.ceil(totalJobs / itemsPerPage));
        setTotalPages(totalPages);

        if (page > totalPages) {
            setCurrentPage(totalPages);
        }
    } catch (error) {
        console.error("Error fetching jobs:", error);
    }
};



    useEffect(() => {
        fetchJobs(currentPage);
    }, [currentPage]);

    const jobsToShow = (jobs && jobs.length > 0) ? jobs.slice((currentPage - 1) * itemsPerPage, currentPage * itemsPerPage) : [];

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
                            {/*<Button onClick={() => setIsModalOpen(true)}*/}
                            {/*        className="bg-indigo-600 mt-9 hover:bg-indigo-700 transition-all duration-300">*/}
                            {/*    Фильтры*/}
                            {/*</Button>*/}
                        </div>
                    </div>
                </div>

                <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8 mt-8">
                    {jobsToShow.map((job) => (
                        <div key={job.Id}
                             className="bg-white p-6 rounded-lg shadow-lg transform hover:scale-105 transition-all duration-300 ease-in-out">
                            <h3 className="text-xl font-semibold text-indigo-600">{job.Title}</h3>
                            <p className="mt-3 text-gray-700">{job.Description}</p>
                            <div className="mt-4 text-gray-400">
                                <span className="block">{job.Location}</span>
                            </div>
                            <Button
                                className="mt-4 w-full bg-indigo-600 hover:bg-indigo-700 transition-all duration-300"
                                href={`/jobs/${job.Id}`}>
                                Просмотреть
                            </Button>
                        </div>
                    ))}
                </div>

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

            <Modal show={isModalOpen} onClose={() => setIsModalOpen(false)}>
                <Modal.Header>Фильтры вакансий</Modal.Header>
                <Modal.Body>
                    <div className="space-y-4">
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
