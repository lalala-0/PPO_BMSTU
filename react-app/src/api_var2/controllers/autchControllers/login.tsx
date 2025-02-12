import api from "../api";

export const login = async (login: string, password: string) => {
  try {
    const response = await api.post("/login", { login, password });
    const token = response.data.token;
    const judgeID = response.data.judge.ID; // Предполагаем, что судья также возвращается с id в ответе

    if (token) {
      sessionStorage.setItem("token", token);
      sessionStorage.setItem("judgeID", judgeID); // Сохраняем id судьи
      window.location.reload(); // Перезагружаем текущую страницу
    }

    return token;
  } catch (error) {
    console.error("Ошибка авторизации", error);
    throw error;
  }
};
