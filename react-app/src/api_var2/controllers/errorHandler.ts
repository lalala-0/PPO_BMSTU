// Функция обработки ошибок
export const handleError = (
    err: any,
    setError: React.Dispatch<React.SetStateAction<string | null>>,
) => {
  if (err && err.response) {
    // Ошибки от сервера (например, 4xx или 5xx)
    const serverMessage = err.response?.data?.message || "Неизвестная ошибка от сервера";
    setError(`Ошибка: ${serverMessage} (код: ${err.response?.status})`);
    alert(`Ошибка: ${serverMessage} (код: ${err.response?.status})`);
  } else if (err && err.request) {
    // Проблемы с сетью
    setError("Ошибка сети. Проверьте подключение к интернету.");
    alert("Ошибка сети. Проверьте подключение к интернету.");
  } else if (err && err.message) {
    // Ошибка при настройке запроса
    setError(`Ошибка запроса: ${err.message}`);
    alert(`Ошибка запроса: ${err.message}`);
  } else {
    // Непредвиденная ошибка
    setError("Произошла неизвестная ошибка.");
    alert("Произошла неизвестная ошибка.");
  }
};
