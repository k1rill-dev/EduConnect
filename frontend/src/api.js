import axios from 'axios';

const apiClient = axios.create({
    baseURL: 'http://localhost:8082/api',
    headers: {
        'Content-Type': 'application/json',
    },
});

let accessToken = null;
let refreshToken = null;

const setAccessToken = (token) => {
    accessToken = token;
    apiClient.defaults.headers.common['Authorization'] = `Bearer ${token}`;
};

// const refreshAccessToken = async () => {
//     try {
//         const response = await axios.post('http://localhost:8082/api/auth/rotate', {
//             RefreshToken: refreshToken,
//         });
//         const {accessToken: newAccessToken, refreshToken: newRefreshToken} = response.data;
//
//         setAccessToken(newAccessToken);
//         refreshToken = newRefreshToken;
//         localStorage.setItem('accessToken', newAccessToken);
//         localStorage.setItem('refreshToken', newRefreshToken);
//
//         return newAccessToken;
//     } catch (error) {
//         console.error('Ошибка обновления токенов:', error);
//         throw error;
//     }
// };
//
// apiClient.interceptors.request.use(
//     (config) => {
//         if (accessToken) {
//             config.headers['Authorization'] = `Bearer ${accessToken}`;
//         }
//         return config;
//     },
//     (error) => Promise.reject(error)
// );
//
// apiClient.interceptors.response.use(
//     (response) => response,
//     async (error) => {
//         const originalRequest = error.config;
//         if (error.response && error.response.status === 401 && !originalRequest._retry) {
//             originalRequest._retry = true;
//             try {
//                 const newAccessToken = await refreshAccessToken();
//                 originalRequest.headers['Authorization'] = `Bearer ${newAccessToken}`;
//                 return apiClient(originalRequest);
//             } catch (refreshError) {
//                 console.error('Не удалось обновить токен:', refreshError);
//                 return Promise.reject(refreshError);
//             }
//         }
//
//         return Promise.reject(error);
//     }
// );

export {apiClient, setAccessToken, refreshToken};
