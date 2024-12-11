import axios from 'axios';

const apiClient = axios.create({
  baseURL: 'https://example.com',
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

const refreshAccessToken = async () => {
  try {
    const response = await axios.post('https://example.com/api/auth/rotate', {
      refreshToken: refreshToken,
    });
    const { accessToken: newAccessToken, refreshToken: newRefreshToken } = response.data;

    setAccessToken(newAccessToken);
    refreshToken = newRefreshToken;

    return newAccessToken;
  } catch (error) {
    console.error('Ошибка обновления токенов:', error);
    throw error;
  }
};

// Перехват запросов
apiClient.interceptors.request.use(
  (config) => {
    if (accessToken) {
      config.headers['Authorization'] = `Bearer ${accessToken}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// Перехват ответов для обработки 401 ошибки
apiClient.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;

    // Если ошибка 401 и запрос ещё не повторялся
    if (error.response && error.response.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;
      try {
        const newAccessToken = await refreshAccessToken();
        originalRequest.headers['Authorization'] = `Bearer ${newAccessToken}`;
        return apiClient(originalRequest);
      } catch (refreshError) {
        console.error('Не удалось обновить токен:', refreshError);
        return Promise.reject(refreshError);
      }
    }

    return Promise.reject(error);
  }
);

export { apiClient, setAccessToken, refreshToken };
