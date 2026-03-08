import axios from "axios";

const rawApiUrl = import.meta.env.VITE_API_URL || "http://localhost:8080/api";
const normalizeApiBaseUrl = (url) => {
  if (!url) return url;
  // Ensure protocol; default to https if missing
  const hasProtocol = /^https?:\/\//i.test(url);
  let normalized = hasProtocol ? url : `https://${url}`;
  // Remove trailing slashes
  normalized = normalized.replace(/\/+$/, "");
  // Ensure the base ends with /api
  if (!/\/api$/i.test(normalized)) {
    normalized = `${normalized}/api`;
  }
  return normalized;
};
const API_BASE_URL = normalizeApiBaseUrl(rawApiUrl);

// Create axios instance
const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

// Request interceptor to add auth token
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem("token");
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  },
);

// Response interceptor to handle errors
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      // Token expired or invalid
      localStorage.removeItem("token");
      localStorage.removeItem("user");
      window.location.href = "/login";
    }
    return Promise.reject(error);
  },
);

// Auth API
export const authAPI = {
  register: async (email, username, password) => {
    const response = await api.post("/register", { email, username, password });
    return response.data;
  },

  login: async (email, password) => {
    const response = await api.post("/login", { email, password });
    if (response.data.token) {
      localStorage.setItem("token", response.data.token);
      localStorage.setItem("user", JSON.stringify(response.data.user));
    }
    return response.data;
  },

  logout: () => {
    localStorage.removeItem("token");
    localStorage.removeItem("user");
  },

  getCurrentUser: () => {
    const userStr = localStorage.getItem("user");
    return userStr ? JSON.parse(userStr) : null;
  },

  isAuthenticated: () => {
    return !!localStorage.getItem("token");
  },
};

// Patterns API
export const patternsAPI = {
  getAll: async () => {
    const response = await api.get("/patterns");
    return response.data;
  },
};

// Problems API
export const problemsAPI = {
  getAll: async (patternId = null) => {
    const params = patternId ? { pattern_id: patternId } : {};
    const response = await api.get("/problems", { params });
    return response.data;
  },
};

// Progress API
export const progressAPI = {
  getAll: async () => {
    const response = await api.get("/progress");
    return response.data;
  },

  update: async (
    problemNumber,
    patternId,
    completed,
    difficulty = "",
    notes = "",
  ) => {
    const response = await api.post("/progress", {
      problem_number: problemNumber,
      pattern_id: patternId,
      completed,
      difficulty,
      notes,
    });
    return response.data;
  },

  getStats: async () => {
    const response = await api.get("/stats");
    return response.data;
  },
};

export default api;
