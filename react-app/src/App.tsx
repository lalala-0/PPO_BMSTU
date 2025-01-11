import React from "react";
import { BrowserRouter as Router, Route, Routes, Link } from "react-router-dom";
import { GetRatingsView, RatingView } from "./ui";
import { CrewView, RaceView } from "./ui";
import { ParticipantView } from "./ui";
// import { GetRatingsView as GetRatingsViewAPI,
//         GetRatingView as GetRatingViewAPI} from './api';

const App: React.FC = () => {
  return (
    <Router>
      {/* Хедер с навигацией */}
      <header style={headerStyle}>
        <nav>
          <Link to="/" style={linkStyle}>
            Главная
          </Link>
          <Link to="/ratings" style={linkStyle}>
            Список рейтингов
          </Link>
        </nav>
      </header>

      <Routes>
        <Route path="/" element={<div>Добро пожаловать!</div>} />
        <Route
          path="/participants/:participantID"
          element={<ParticipantView />}
        />
        <Route path="/ratings" element={<GetRatingsView />} />
        <Route path="/ratings/:ratingID" element={<RatingView />} />
        <Route path="/ratings/:ratingID/races/:raceID" element={<RaceView />} />
        <Route path="/ratings/:ratingID/crews/:crewID" element={<CrewView />} />

        {/*<Route path="/api/ratings" element={<GetRatingsViewAPI />} />*/}
        {/*<Route path="/api/ratings/:ratingID" element={<GetRatingViewAPI />} />*/}
      </Routes>

      {/* Нижний элемент (footer) */}
      <footer style={footerStyle}>
        <p>&copy; 2025, Все права защищены.</p>
      </footer>
    </Router>
  );
};

export default App;

// Стили для хедера
const headerStyle: React.CSSProperties = {
  backgroundColor: "#D9D9D9",
  color: "white",
  textAlign: "center",
  padding: "10px 0",
  position: "sticky",
  top: 0,
  width: "100%",
  zIndex: 1000,
};

// Стили для ссылок в навигации
const linkStyle: React.CSSProperties = {
  color: "#000000",
  textDecoration: "none",
  margin: "0 15px",
  fontSize: "50px",
  fontFamily: "Spectral SC",
};

const footerStyle: React.CSSProperties = {
  backgroundColor: "#C5D1D7",
  color: "#000000",
  fontFamily: "Spectral SC",
  textAlign: "center",
  padding: "40px 0",
  bottom: 0,
  width: "100%",
};
