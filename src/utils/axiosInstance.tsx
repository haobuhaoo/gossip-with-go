import axios, { AxiosError } from "axios";

/**
 * Creates a predefined axios instance.
 */
const axiosInstance = axios.create({
    baseURL: import.meta.env.VITE_API_BASE,
    timeout: 10000,
    headers: { "Content-Type": "application/json" },
})

/**
 * Intercepts every outgoing request and attaches JWT token if present to Authorization header.
 */
axiosInstance.interceptors.request.use(
    (config) => {
        const token: string | null = localStorage.getItem("token");
        if (token) {
            config.headers = config.headers ?? {}
            config.headers["Authorization"] = `Bearer ${token}`;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
)

/**
 * Intercepts every response received and returns successful responses or normalises error
 * for failed requests.
 */
axiosInstance.interceptors.response.use(
    (response) => response,
    (error: AxiosError<any>) => {
        if (!error.response) {
            return Promise.reject("Network Error");
        }

        const msg: string = error.response.data?.messages?.[0] ?? "Request failed";
        return Promise.reject(msg);
    }
)

export default axiosInstance;
