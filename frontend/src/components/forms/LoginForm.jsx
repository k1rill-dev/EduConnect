import React, {useState} from 'react';
import axios from "axios";
import {useNavigate} from "react-router-dom";

// const login = async (sendData) => {
//     const {data, status} = await api.post('/api/login', sendData, {
//         headers: {
//             'Content-Type': 'application/json',
//             Accept: 'application/json'
//         },
//         withCredentials: true
//     });
//     return data;
//
// }

const LoginForm = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const nav = useNavigate();
    const handleLogin = async (event) => {
        event.preventDefault();
        let sendData = {
            email: email,
            password: password
        }
        // let data = await login(sendData).then(res => res);
        // localStorage.setItem('data', JSON.stringify(data));
        // nav('/')
    }
    return (
        <div className="flex items-center justify-center min-h-screen bg-gradient-to-b from-white to-blue-400">
            <div className="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
                <h2 className="text-3xl font-bold mb-6 text-gray-900 text-center">Войти</h2>
                <form onSubmit={handleLogin} className="space-y-4">
                    <div>
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="email">
                            Email
                        </label>
                        <input
                            onChange={(event) => setEmail(event.target.value)}
                            type="email"
                            id="email"
                            placeholder="Введите ваш email"
                            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                        />
                    </div>
                    <div>
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="password">
                            Пароль
                        </label>
                        <input
                            onChange={(event) => setPassword(event.target.value)}
                            type="password"
                            id="password"
                            placeholder="Введите ваш пароль"
                            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                        />
                    </div>
                    <div>
                        <button
                            type="submit"
                            className="w-full bg-indigo-600 hover:bg-indigo-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                        >
                            Войти
                        </button>
                    </div>
                </form>
                <div className="mt-4 text-center">
                    <p className="text-sm text-gray-600">Нет аккаунта? <a href="/register"
                                                                          className="font-medium text-indigo-600 hover:text-indigo-500">Зарегистрироваться</a>
                    </p>
                </div>
            </div>
        </div>
    );
};

export default LoginForm;