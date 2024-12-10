import React, { useState } from "react";
import { Button, Label, TextInput, FileInput, Textarea } from "flowbite-react";

const EditProfilePage = () => {
  const [userData, setUserData] = useState({
    username: "John Doe",
    email: "john.doe@example.com",
    role: "student",
    bio: "This is my bio.",
    profilePicture: null,
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setUserData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleProfilePictureChange = (e) => {
    setUserData((prev) => ({
      ...prev,
      profilePicture: e.target.files[0],
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log("Updated user data:", userData);
    alert("Profile updated successfully!");
  };

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="bg-gradient-to-r from-blue-500 to-purple-600 text-white py-10 shadow-md">
        <div className="container mx-auto px-6 text-center">
          <h1 className="text-4xl font-bold">Редактировать профиль</h1>
          <p className="mt-2 text-lg">Обновите свои данные для профиля.</p>
        </div>
      </div>

      <div className="container mx-auto px-6 py-10">
        <form onSubmit={handleSubmit} className="bg-white p-6 rounded-lg shadow-lg space-y-6">
          <div className="flex justify-center">
            <div className="w-24 h-24 rounded-full overflow-hidden bg-gray-200">
              <img
                src={userData.profilePicture ? URL.createObjectURL(userData.profilePicture) : "/default-profile.png"}
                alt="Profile"
                className="w-full h-full object-cover"
              />
            </div>
          </div>

          <div>
            <Label htmlFor="username" className="mb-2 block">Имя</Label>
            <TextInput
              id="username"
              name="username"
              type="text"
              placeholder="Введите ваше имя"
              value={userData.username}
              onChange={handleChange}
              required
            />
          </div>

          <div>
            <Label htmlFor="email" className="mb-2 block">Электронная почта</Label>
            <TextInput
              id="email"
              name="email"
              type="email"
              placeholder="Введите вашу почту"
              value={userData.email}
              onChange={handleChange}
              required
            />
          </div>

          <div>
            <Label htmlFor="role" className="mb-2 block">Роль</Label>
            <select
              id="role"
              name="role"
              value={userData.role}
              onChange={handleChange}
              className="w-full p-2 rounded-md border border-gray-300 focus:ring-2 focus:ring-blue-500"
            >
              <option value="student">Студент</option>
              <option value="teacher">Преподаватель</option>
              <option value="company">Компания</option>
            </select>
          </div>

          <div>
            <Label htmlFor="bio" className="mb-2 block">О себе</Label>
            <Textarea
              id="bio"
              name="bio"
              placeholder="Напишите о себе"
              value={userData.bio}
              onChange={handleChange}
              rows={4}
            />
          </div>

          <div>
            <Label htmlFor="profilePicture" className="mb-2 block">Фотография профиля</Label>
            <FileInput
              id="profilePicture"
              name="profilePicture"
              onChange={handleProfilePictureChange}
            />
          </div>

          <div className="flex justify-end">
            <Button className="bg-indigo-600 hover:bg-indigo-700" type="submit">
              Сохранить изменения
            </Button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default EditProfilePage;
