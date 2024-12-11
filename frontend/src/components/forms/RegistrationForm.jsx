import React from 'react';
import {Button} from 'flowbite-react';
import {useNavigate} from 'react-router-dom';
import {apiClient} from "../../api";

const register = async (user) => {
    try {
        const {data} = await apiClient.post('/auth/sign-up', user, {
            headers: {
                'Content-Type': 'application/json',
            },
        });
        return data;
    } catch (error) {
        console.error('Error during registration:', error);
        throw error;
    }
};

const RegistrationForm = () => {
    const [firstName, setFirstName] = React.useState('');
    const [lastName, setLastName] = React.useState('');
    const [email, setEmail] = React.useState('');
    const [password, setPassword] = React.useState('');
    const [repeatPassword, setRepeatPassword] = React.useState('');
    const [bio, setBio] = React.useState('');
    const [profilePicture, setProfilePicture] = React.useState('');
    const [role, setRole] = React.useState('student'); // Default role: 'student'
    const [errorMessage, setErrorMessage] = React.useState('');
    const [passwordVisible, setPasswordVisible] = React.useState(false);
    const [repeatPasswordVisible, setRepeatPasswordVisible] = React.useState(false);
    const nav = useNavigate();

    const handleFileChange = (event) => {
        const file = event.target.files[0];
        if (file) {
            const reader = new FileReader();
            reader.onload = () => {
                setProfilePicture(reader.result.split(",")[1]);
            };
            reader.readAsDataURL(file);
        }
    };

    const handleSendUserData = async (event) => {
        event.preventDefault();
        const sendData = {
            first_name: firstName,
            surname: lastName,
            email: email,
            password: password,
            role: role,
            bio: bio,
            picture: profilePicture,
        };

        try {
            console.log(sendData)
            const response = await register(sendData);
            console.log(response)
            localStorage.setItem('accessToken', response.access_token);
            localStorage.setItem('refreshToken', response.refresh_token);
            localStorage.setItem('userData', JSON.stringify(response));
            nav('/');
        } catch (error) {
            setErrorMessage('Произошла ошибка при регистрации');
        }
    };

    const togglePasswordVisibility = () => setPasswordVisible(!passwordVisible);
    const toggleRepeatPasswordVisibility = () => setRepeatPasswordVisible(!repeatPasswordVisible);

    return (
        <div className="min-h-screen bg-gradient-to-r from-indigo-50 to-purple-100 flex items-center justify-center">
            <div className="bg-white p-8 rounded-lg shadow-lg w-full max-w-md">
                <h2 className="text-3xl font-extrabold text-gray-800 mb-6 text-center">Регистрация</h2>
                <form onSubmit={handleSendUserData} className="space-y-6">
                    <div className="grid grid-cols-2 gap-6">
                        <div>
                            <label className="block text-gray-700 text-sm font-semibold mb-2"
                                   htmlFor="firstName">Имя</label>
                            <input
                                onChange={(e) => setFirstName(e.target.value)}
                                type="text"
                                id="firstName"
                                placeholder="Введите ваше имя"
                                className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-600"
                            />
                        </div>
                        <div>
                            <label className="block text-gray-700 text-sm font-semibold mb-2"
                                   htmlFor="lastName">Фамилия</label>
                            <input
                                onChange={(e) => setLastName(e.target.value)}
                                type="text"
                                id="lastName"
                                placeholder="Введите вашу фамилию"
                                className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-600"
                            />
                        </div>
                    </div>
                    <div>
                        <label className="block text-gray-700 text-sm font-semibold mb-2" htmlFor="email">Email</label>
                        <input
                            onChange={(e) => setEmail(e.target.value)}
                            type="email"
                            id="email"
                            placeholder="Введите ваш email"
                            className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-600"
                        />
                    </div>
                    <div>
                        <label className="block text-gray-700 text-sm font-semibold mb-2" htmlFor="bio">Описание
                            профиля</label>
                        <textarea
                            onChange={(e) => setBio(e.target.value)}
                            id="bio"
                            placeholder="Напишите немного о себе"
                            className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-600"
                        />
                    </div>
                    <div>
                        <label className="block text-gray-700 text-sm font-semibold mb-2" htmlFor="photo">Фото
                            профиля</label>
                        <input
                            onChange={handleFileChange}
                            type="file"
                            id="photo"
                            accept="image/*"
                            className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-600"
                        />
                    </div>
                    <div>
                        <label className="block text-gray-700 text-sm font-semibold mb-2" htmlFor="role">Роль</label>
                        <select
                            id="role"
                            value={role}
                            onChange={(e) => setRole(e.target.value)}
                            className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-600"
                        >
                            <option value="student">Студент</option>
                            <option value="teacher">Учитель</option>
                            <option value="employer">Работодатель</option>
                        </select>
                    </div>
                    <div>
                        <label className="block text-gray-700 text-sm font-semibold mb-2"
                               htmlFor="password">Пароль</label>
                        <div className="relative">
                            <input
                                onChange={(e) => setPassword(e.target.value)}
                                type={passwordVisible ? 'text' : 'password'}
                                id="password"
                                placeholder="Введите ваш пароль"
                                className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-600"
                            />
                            <button
                                type="button"
                                onClick={togglePasswordVisibility}
                                className="absolute top-1/2 right-3 transform -translate-y-1/2 text-gray-500"
                            >
                                {passwordVisible ? 'Скрыть' : 'Показать'}
                            </button>
                        </div>
                    </div>
                    <div>
                        <label className="block text-gray-700 text-sm font-semibold mb-2" htmlFor="confirm-password">Повторите
                            пароль</label>
                        <div className="relative">
                            <input
                                onChange={(e) => setRepeatPassword(e.target.value)}
                                type={repeatPasswordVisible ? 'text' : 'password'}
                                id="confirm-password"
                                placeholder="Повторите ваш пароль"
                                className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-600"
                            />
                            <button
                                type="button"
                                onClick={toggleRepeatPasswordVisibility}
                                className="absolute top-1/2 right-3 transform -translate-y-1/2 text-gray-500"
                            >
                                {repeatPasswordVisible ? 'Скрыть' : 'Показать'}
                            </button>
                        </div>
                    </div>
                    <div>
                        <Button
                            type="submit"
                            className="w-full bg-indigo-600 hover:bg-indigo-700 text-white py-3 rounded-lg"
                        >
                            Зарегистрироваться
                        </Button>
                    </div>
                </form>
                {errorMessage && (
                    <div className="text-red-500 text-sm font-medium mt-2">{errorMessage}</div>
                )}
            </div>
        </div>
    );
};

export default RegistrationForm;
