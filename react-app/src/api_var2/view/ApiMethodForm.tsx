import React, { useState } from "react";

interface ApiMethodFormProps {
  method: string;
  url: string;
  parameters: any[];
  onSubmit: (params: any) => void; // Обработчик отправки данных
}

const ApiMethodForm: React.FC<ApiMethodFormProps> = ({
  method,
  url,
  parameters,
  onSubmit,
}) => {
  const [formData, setFormData] = useState<any>({});

  // Обработчик изменения значения в полях формы
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  // Обработчик отправки формы
  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit(formData); // Передаем данные в родительский компонент (контроллер)
  };

  return (
    <form onSubmit={handleSubmit}>
      <h3>
        {method.toUpperCase()} {url}
      </h3>
      {/* Перебираем параметры и создаем поля ввода */}
      {parameters.map((param: any) => (
        <div key={param.name}>
          <label>{param.name}</label>
          <input
            type="text"
            name={param.name}
            onChange={handleChange}
            placeholder={param.name}
          />
        </div>
      ))}
      <button type="submit">Submit</button>
    </form>
  );
};

export default ApiMethodForm;
