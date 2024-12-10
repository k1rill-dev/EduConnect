import React, { useState } from 'react';
import { Button } from 'flowbite-react';
import { useNavigate } from 'react-router-dom';

const LoginForm = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const nav = useNavigate();

  const handleLogin = async (event) => {
    event.preventDefault();
    let sendData = {
      email: email,
      password: password,
    };
    // let data = await login(sendData).then(res => res);
    // localStorage.setItem('data', JSON.stringify(data));
    // nav('/');
  };

  return (
    <div className="min-h-screen bg-gradient-to-r from-indigo-50 to-purple-100 flex items-center justify-center">
      <div className="bg-white p-8 rounded-lg shadow-lg w-full max-w-md">
        <h2 className="text-3xl font-extrabold text-gray-800 mb-6 text-center">Войти</h2>
        <form onSubmit={handleLogin} className="space-y-6">
          <div>
            <label className="block text-gray-700 text-sm font-semibold mb-2" htmlFor="email">
              Email
            </label>
            <input
              onChange={(event) => setEmail(event.target.value)}
              type="email"
              id="email"
              placeholder="Введите ваш email"
              className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-600"
            />
          </div>
          <div>
            <label className="block text-gray-700 text-sm font-semibold mb-2" htmlFor="password">
              Пароль
            </label>
            <input
              onChange={(event) => setPassword(event.target.value)}
              type="password"
              id="password"
              placeholder="Введите ваш пароль"
              className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-600"
            />
          </div>
          <div>
            <Button type="submit" className="w-full bg-indigo-600 hover:bg-indigo-700 text-white py-3 rounded-lg">
              Войти
            </Button>
          </div>
        </form>
        <div className="mt-4 text-center">
          <p className="text-sm text-gray-600">
            Нет аккаунта?{' '}
            <a href="/register" className="font-medium text-indigo-600 hover:text-indigo-500">
              Зарегистрироваться
            </a>
          </p>
        </div>
      </div>
    </div>
  );
};

export default LoginForm;
