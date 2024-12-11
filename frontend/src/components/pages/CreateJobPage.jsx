import React, { useState } from 'react';
import axios from 'axios';
import { Button, Label, TextInput, Textarea } from 'flowbite-react';
import {apiClient} from "../../api";

const CreateJobPage = () => {
  const [formData, setFormData] = useState({
    title: '',
    description: '',
    location: '',
  });

  const [isSubmitting, setIsSubmitting] = useState(false);
   const userProfile = JSON.parse(localStorage.getItem('userData')) || {};
  const userId = userProfile.id;

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsSubmitting(true);

    try {
      const response = await apiClient.post('/jobs', formData, {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${userProfile.access_token}`,
        },
      });

      console.log('Создана вакансия:', response.data);
      alert('Вакансия успешно создана!');
      setFormData({ title: '', description: '', location: '' });
    } catch (error) {
      console.error('Ошибка при создании вакансии:', error);
      alert('Произошла ошибка при создании вакансии.');
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="bg-gradient-to-r from-purple-500 to-indigo-600 text-white py-10 shadow-md">
        <div className="container mx-auto px-6 text-center">
          <h1 className="text-4xl font-bold">Создать вакансию</h1>
          <p className="mt-2 text-lg">
            Заполните форму ниже, чтобы опубликовать вакансию.
          </p>
        </div>
      </div>

      <div className="container mx-auto px-6 py-10">
        <form
          onSubmit={handleSubmit}
          className="bg-white p-6 rounded-lg shadow-lg space-y-6"
        >
          <div>
            <Label htmlFor="title" className="mb-2 block">
              Название вакансии
            </Label>
            <TextInput
              id="title"
              name="title"
              type="text"
              placeholder="Введите название вакансии"
              value={formData.title}
              onChange={handleChange}
              required
            />
          </div>

          <div>
            <Label htmlFor="description" className="mb-2 block">
              Описание вакансии
            </Label>
            <Textarea
              id="description"
              name="description"
              placeholder="Опишите основные требования и обязанности"
              rows={4}
              value={formData.description}
              onChange={handleChange}
              required
            />
          </div>

          <div>
            <Label htmlFor="location" className="mb-2 block">
              Местоположение
            </Label>
            <TextInput
              id="location"
              name="location"
              type="text"
              placeholder="Введите местоположение или 'удаленно'"
              value={formData.location}
              onChange={handleChange}
              required
            />
          </div>

          <div className="flex justify-end">
            <Button
              type="submit"
              className="bg-indigo-600 hover:bg-indigo-700"
              disabled={isSubmitting}
            >
              {isSubmitting ? 'Отправка...' : 'Опубликовать вакансию'}
            </Button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default CreateJobPage;
