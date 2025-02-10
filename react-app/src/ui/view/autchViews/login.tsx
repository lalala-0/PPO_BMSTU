import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { login } from "../../controllers/autchControllers/login";

const Login: React.FC = () => {
  const [username, setUsername] = useState(""); // Логин
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    try {
      const token = await login(username, password);
      if (token) {
        navigate("/dashboard"); // Перенаправление после успешного входа
      }
    } catch {
      setError("Неверные данные для входа");
    }
  };

  return (
    <div className="flex items-center justify-center h-screen">
      <form
        onSubmit={handleSubmit}
        className="bg-white p-6 rounded-lg shadow-lg w-80"
      >
        <h2 className="text-xl font-semibold mb-4 text-center">Вход</h2>

        {error && <p className="text-red-500 text-center">{error}</p>}

        <div className="mb-4">
          <label className="block text-sm font-medium">Логин</label>
          <input
            type="text"
            className="w-full p-2 border rounded"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
        </div>

        <div className="mb-4">
          <label className="block text-sm font-medium">Пароль</label>
          <input
            type="text"
            className="w-full p-2 border rounded"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>

        <button
          type="submit"
          className="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700 transition"
        >
          Войти
        </button>
      </form>
    </div>
  );
};

export default Login;
