import React from "react";

interface FiltersProps {
  filters: { [key: string]: string };
  onFilterChange: (key: string, value: string) => void;
}

const Filters: React.FC<FiltersProps> = ({ filters, onFilterChange }) => {
  return (
    <div style={{ display: "flex", marginBottom: "20px" }}>
      {Object.entries(filters).map(([key, value]) => (
        <div key={key} style={{ marginRight: "20px", width: "25%" }}>
          <input
            type="text"
            placeholder={`Поиск по ${key}`}
            value={value}
            onChange={(e) => onFilterChange(key, e.target.value)}
            style={{ width: "100%" }}
          />
        </div>
      ))}
    </div>
  );
};

export default Filters;
