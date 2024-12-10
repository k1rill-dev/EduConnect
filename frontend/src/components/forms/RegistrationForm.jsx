import React from 'react';
import { Button } from 'flowbite-react';
import { useNavigate } from 'react-router-dom';

const register = async (user) => {
  // const { data, status } = await api.post("/api/register", user, {
  //     headers: {
  //         'Content-Type': 'application/json'
  //     },
  //     withCredentials: true
  // });
  // return data
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

  const handleEmailChange = (event) => {
    if (
      !event.target.value ||
      !/^(([^<>()[\]\.,;:\s@\"]+(\.[^<>()[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;@\s]{2,})$/i.test(event.target.value)
    ) {
      setErrorMessage('Неправильный формат email!');
    } else {
      setErrorMessage('');
      setEmail(event.target.value);
    }
  };

  const handlePasswordChange = (event) => {
    if (
      !event.target.value ||
      !/^(?=.*\d)(?=.*[!@#$%^&*])(?=.*[a-z])(?=.*[A-Z]).{8,}$/.test(event.target.value)
    ) {
      setErrorMessage(
        'Пароль слишком слабый! Пароль должен содержать строчные и заглавные буквы, быть не менее 8 символов, а также содержать спец. символы'
      );
    } else {
      setErrorMessage('');
      setPassword(event.target.value);
    }
  };

  const handleRepeatPassword = (event) => {
    if (password !== event.target.value) {
      setErrorMessage('Пароли не совпадают!');
    } else {
      setErrorMessage('');
    }
    setRepeatPassword(event.target.value);
  };

  const handleSendUserData = async (event) => {
    event.preventDefault();
    let sendData = {
      first_name: firstName,
      last_name: lastName,
      email: email,
      password_hash: password, // You should hash the password before sending it to the server
      role: role, // 'teacher', 'student', or 'employer'
      bio: bio,
      profile_picture: profilePicture,
      is_active: true,
      is_superuser: false,
    };
    let response = await register(sendData);
    if (response.success) {
      nav('/login');
    }
  };

  const togglePasswordVisibility = () => {
    setPasswordVisible(!passwordVisible);
  };

  const toggleRepeatPasswordVisibility = () => {
    setRepeatPasswordVisible(!repeatPasswordVisible);
  };

  return (
    <div className="min-h-screen bg-gradient-to-r from-indigo-50 to-purple-100 flex items-center justify-center">
      <div className="bg-white p-8 rounded-lg shadow-lg w-full max-w-md">
        <h2 className="text-3xl font-extrabold text-gray-800 mb-6 text-center">Регистрация</h2>
        <form onSubmit={handleSendUserData} className="space-y-6">
          <div className="grid grid-cols-2 gap-6">
            <div>
              <label className="block text-gray-700 text-sm font-semibold mb-2" htmlFor="firstName">
                Имя
              </label>
              <input
                  onChange={(e) => setFirstName(e.target.value)}
                  type="text"
                  id="firstName"
                  placeholder="Введите ваше имя"
                  className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-600"
              />
            </div>
            <div>
              <label className="block text-gray-700 text-sm font-semibold mb-2" htmlFor="lastName">
                Фамилия
              </label>
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
            <label className="block text-gray-700 text-sm font-semibold mb-2" htmlFor="email">
              Email
            </label>
            <input
                onChange={handleEmailChange}
                type="email"
                id="email"
                placeholder="Введите ваш email"
                className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-600"
            />
          </div>
          <div>
            <label className="block text-gray-700 text-sm font-semibold mb-2" htmlFor="bio">
              Описание профиля
            </label>
            <textarea
                onChange={(e) => setBio(e.target.value)}
                id="bio"
                placeholder="Напишите немного о себе"
                className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-600"
            />
          </div>
          <div>
            <label className="block text-gray-700 text-sm font-semibold mb-2" htmlFor="photo">
              Фото профиля
            </label>
            <input
                onChange={(e) => setProfilePicture(e.target.value)}
                type="text"
                id="photo"
                placeholder="Введите ссылку на фото профиля"
                className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-600"
            />
          </div>
          <div>
            <label className="block text-gray-700 text-sm font-semibold mb-2" htmlFor="role">
              Роль
            </label>
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
            <label className="block text-gray-700 text-sm font-semibold mb-2" htmlFor="password">
              Пароль
            </label>
            <div className="relative">
              <input
                  onChange={handlePasswordChange}
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
            <label className="block text-gray-700 text-sm font-semibold mb-2" htmlFor="confirm-password">
              Повторите пароль
            </label>
            <div className="relative">
              <input
                  onChange={handleRepeatPassword}
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
        <div className="mt-4 text-center">
          <p className="text-sm text-gray-600">
            Уже есть аккаунт?{' '}
            <a href="/login" className="font-medium text-indigo-600 hover:text-indigo-500">
              Войдите
            </a>
          </p>
        </div>
      </div>
    </div>
  );
};

export default RegistrationForm;
