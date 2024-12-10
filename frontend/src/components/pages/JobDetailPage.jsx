import React from "react";
import { useParams } from "react-router-dom";
import {Button} from "flowbite-react"; // Для получения ID вакансии из URL

const JobDetailPage = () => {

  const { jobId } = useParams();
  console.log(jobId)
  // Пример вакансий
  const jobs = [
    { jobId: 1, title: "Software Engineer", company: "TechCorp", location: "New York", description: "Develop and maintain web applications." },
    { jobId: 2, title: "Frontend Developer", company: "WebSolutions", location: "San Francisco", description: "Work with React and Tailwind CSS." },
    { jobId: 3, title: "Backend Developer", company: "DevHub", location: "Chicago", description: "Build and optimize APIs." },
  ];

  const job = jobs.find((job) => job.jobId === parseInt(jobId));

  if (!job) {
    return <p>Вакансия не найдена!</p>;
  }

  return (
    <div className="min-h-screen bg-gradient-to-r from-indigo-50 to-purple-100">
      <div className="bg-gradient-to-r from-blue-600 to-purple-600 text-white py-12 shadow-md">
        <div className="container mx-auto px-6 text-center">
          <h1 className="text-4xl font-extrabold">{job.title}</h1>
          <p className="mt-2 text-lg">{job.company}</p>
        </div>
      </div>

      <div className="container mx-auto px-6 py-12">
        <div className="bg-white p-8 rounded-lg shadow-xl space-y-6">
          <h2 className="text-xl font-semibold text-indigo-600">Локация:</h2>
          <h3 className="mt-3 text-lg font-medium text-gray-800">{job.location}</h3>
          <h2 className="text-xl font-semibold text-indigo-600">Описание</h2>
          <p>{job.description}</p>
          <Button className="mt-6 bg-indigo-600 hover:bg-indigo-700 transition-all duration-300">
            Подать заявку
          </Button>
        </div>
      </div>
    </div>
  );
};

export default JobDetailPage;
