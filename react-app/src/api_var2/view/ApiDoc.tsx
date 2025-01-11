// src/components/ApiDoc.tsx

import React, { useState, useEffect } from "react";
import { fetchApiData, makeApiRequest } from "../controllers/ApiController"; // Импортируем контроллер
import ApiMethodForm from "./ApiMethodForm";

const ApiDoc: React.FC = () => {
  const [apiData, setApiData] = useState<any>(null);
  const [response, setResponse] = useState<any>(null);
  const [error, setError] = useState<string>("");

  // Загружаем документацию API
  useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await fetchApiData(); // Получаем данные через контроллер
        setApiData(data);
      } catch (error) {
        console.error("Error loading API doc:", error);
      }
    };
    fetchData();
  }, []);

  // Обработчик отправки формы
  const handleApiSubmit = async (
    formData: any,
    method: string,
    url: string,
  ) => {
    try {
      const result = await makeApiRequest(method, url, formData);
      setResponse(result);
    } catch (error) {
      setError("Error executing API request");
    }
  };

  if (!apiData) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <h1>{apiData.info.title}</h1>
      <p>{apiData.info.description}</p>

      {/* Перебираем пути и методы API */}
      {apiData.paths &&
        Object.keys(apiData.paths).map((path) => {
          const methods = apiData.paths[path];

          return Object.keys(methods).map((method) => {
            const apiMethod = methods[method];
            return (
              <div key={`${method}-${path}`} style={{ marginBottom: "20px" }}>
                <ApiMethodForm
                  method={method} // HTTP метод (GET, POST и т.д.)
                  url={path} // URL пути
                  parameters={apiMethod.parameters || []} // Параметры запроса
                  onSubmit={(formData) =>
                    handleApiSubmit(formData, method, path)
                  } // Передаем обработчик
                />
              </div>
            );
          });
        })}

      {/* Отображаем ответ или ошибку */}
      {response && <pre>{JSON.stringify(response, null, 2)}</pre>}
      {error && <p>{error}</p>}
    </div>
  );
};

export default ApiDoc;
