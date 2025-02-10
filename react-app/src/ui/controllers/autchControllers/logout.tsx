// Логаут
export const logout = () => {
  sessionStorage.removeItem("token");
  window.location.reload(); // Перезагружаем текущую страницу
};
