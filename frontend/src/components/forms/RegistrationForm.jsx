import React from 'react';
import {useNavigate} from "react-router-dom";
// import api, {BACKEND_API_URL} from "../../api";
// import {handleYandexLogin} from "../../tools/yandex";


const register = async (user) => {
    // const {data, status} = await api.post("/api/register", user, {
    //     headers: {
    //         'Content-Type': 'application/json'
    //     },
    //     withCredentials: true
    // });
    // return data
}

const RegistrationForm = () => {
    const [isSpecialist, setIsSpecialist] = React.useState(false);
    const [currentStepIndex, setCurrentStepIndex] = React.useState(0);
    const [firstName, setFirstName] = React.useState('');
    const [lastName, setLastName] = React.useState('');
    const [email, setEmail] = React.useState('');
    const [password, setPassword] = React.useState('');
    const [repeatPassword, setRepeatPassword] = React.useState('');
    const [errorMessage, setErrorMessage] = React.useState('');
    const [userId, setUserId] = React.useState('');
    const [speciality, setSpeciality] = React.useState('');
    const [bio, setBio] = React.useState('');
    const nav = useNavigate();
    const handleEmailChange = (event) => {
        if ((!event.target.value) || !(/^(([^<>()[\]\.,;:\s@\"]+(\.[^<>()[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i.test(event.target.value))) {
            setErrorMessage("Неправильный формат email!")
        } else {
            setErrorMessage("")
            setEmail(event.target.value)
        }
    }
    const handlePasswordChange = (event) => {
        if ((!event.target.value) || !(/^(?=.*\d)(?=.*[!@#$%^&*])(?=.*[a-z])(?=.*[A-Z]).{8,}$/.test(event.target.value))) {
            setErrorMessage("Пароль слишком слабый! Пароль должен содержать строчные и заглавные буквы, быть не менее 8 символов, а также содержать спец. символы")
        } else {
            setErrorMessage("");
            setPassword(event.target.value)
        }

    };

    const handleRepeatPassword = (event) => {
        if (password !== event.target.value) {
            setErrorMessage("Пароли не совпадают!")
        } else {
            setErrorMessage("")
        }
        setRepeatPassword(event.target.value);
    }

    const handleNameChange = (event) => {
        setFirstName(event.target.value);
    }

    const handleSurnameChange = (event) => {
        setLastName(event.target.value);

    }
    const handleNextStep = () => {
        setCurrentStepIndex(currentStepIndex + 1);
    }
    const handlePreviousStep = () => {
        if (currentStepIndex === 0) {
            handleNextStep();
        }
        setCurrentStepIndex(currentStepIndex - 1);
    }
    const handleSendUserData = async (event) => {
        event.preventDefault()
        let sendData = {
            first_name: firstName,
            last_name: lastName,
            email: email,
            password: password,
            is_active: true,
            is_superuser: false
        }
        let response = await register(sendData);
        setUserId(response.res)
        if (isSpecialist) {
            handleNextStep();
        } else {
            nav('/login')
        }
    }
    return (
        <div className="flex items-center justify-center min-h-screen bg-gradient-to-b from-white to-blue-400">
            <div className="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
                <h2 className="text-3xl font-bold mb-6 text-gray-900 text-center">Регистрация</h2>
                <form onSubmit={handleSendUserData} className="space-y-4">
                    <div className="grid grid-cols-2 gap-4">
                        <div>
                            <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="firstName">
                                Имя
                            </label>
                            <input
                                onChange={handleNameChange}
                                type="text"
                                id="firstName"
                                placeholder="Введите ваше имя"
                                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                            />
                        </div>
                        <div>
                            <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="lastName">
                                Фамилия
                            </label>
                            <input
                                onChange={handleSurnameChange}
                                type="text"
                                id="lastName"
                                placeholder="Введите вашу фамилию"
                                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                            />
                        </div>
                    </div>
                    <div>
                        <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="email">
                            Email
                        </label>
                        <input
                            onChange={handleEmailChange}
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
                            onChange={handlePasswordChange}
                            type="password"
                            id="password"
                            placeholder="Введите ваш пароль"
                            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                        />
                    </div>
                    <div>
                        <label className="block text-gray-700 text-sm font-bold mb-2"
                               htmlFor="confirm-password">
                            Повторите пароль
                        </label>
                        <input
                            onChange={handleRepeatPassword}
                            type="password"
                            id="confirm-password"
                            placeholder="Повторите ваш пароль"
                            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                        />
                    </div>
                    <div>
                        <button
                            type="submit"
                            className="w-full bg-indigo-600 hover:bg-indigo-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                        >
                            Зарегистрироваться
                        </button>
                    </div>
                </form>
                {errorMessage && (
                    <div className="text-red-500 text-sm font-medium">{errorMessage}</div>
                )}
                <div className="mt-4 text-center">
                    <p className="text-sm text-gray-600">Уже есть аккаунт? <a href="/login"
                                                                              className="font-medium text-indigo-600 hover:text-indigo-500">Войдите</a>
                    </p>
                </div>
            </div>
        </div>
    );

};

export default RegistrationForm;