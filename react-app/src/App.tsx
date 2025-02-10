import React, { useEffect } from "react";
import { BrowserRouter as Router, Route, Routes, Link } from "react-router-dom";
import { GetRatingsView, RatingView } from "./ui";
import { CrewView, RaceView } from "./ui";
import { ParticipantView } from "./ui";
import StartProcedure from "./ui/view/raceViews/startView";
import FinishProcedure from "./ui/view/raceViews/finishView";
import ProtestView from "./ui/view/protestViews/protestView";
import Login from "./ui/view/autchViews/login";
import { isAuthenticated } from "./ui/controllers/autchControllers/isAutchenticated";
import { logout } from "./ui/controllers/autchControllers/logout";
import JudgeProfile from "./ui/view/judgeViews/profile";
import JudgeDashboard from "./ui/view/judgeViews/dashboard";

const Footer: React.FC = () => {
  return (
    <footer style={footerStyle}>
      {isAuthenticated() ? (
        <>
          <Link to="/judgeDashboard" style={bottonLinkStyle}>
            Панель управления
          </Link>
          <Link to="/profile" style={bottonLinkStyle}>
            Профиль
          </Link>
          <button
            onClick={logout}
            style={{
              ...bottonLinkStyle,
              border: "none",
              background: "none",
              cursor: "pointer",
            }}
          >
            Выйти
          </button>
        </>
      ) : (
        <Link to="/login" style={bottonLinkStyle}>
          Войти
        </Link>
      )}
      <p>&copy; 2025, Все права защищены.</p>
    </footer>
  );
};

const App: React.FC = () => {
  useEffect(() => {
    // Обновляем класс "authenticated" при загрузке приложения
    document.body.classList.toggle("authenticated", isAuthenticated());
  }, []);
  return (
    <Router>
      {/* Хедер с навигацией */}
      <header style={headerStyle}>
        <nav>
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
        <Route path="/profile" element={<JudgeProfile />} />
        <Route path="/ratings" element={<GetRatingsView />} />
        <Route path="/ratings/:ratingID" element={<RatingView />} />
        <Route path="/ratings/:ratingID/races/:raceID" element={<RaceView />} />
        <Route
          path="/ratings/:ratingID/races/:raceID/startProcedure"
          element={<StartProcedure />}
        />
        <Route
          path="/ratings/:ratingID/races/:raceID/finishProcedure"
          element={<FinishProcedure />}
        />
        <Route
          path="/ratings/:ratingID/races/:raceID/protests/:protestID"
          element={<ProtestView />}
        />
        <Route path="/ratings/:ratingID/crews/:crewID" element={<CrewView />} />
        <Route path="/login" element={<Login />} />
        <Route path="/judgeDashboard" element={<JudgeDashboard />} />

        {/*<Route path="/api/ratings" element={<GetRatingsViewAPI />} />*/}
        {/*<Route path="/api/ratings/:ratingID" element={<GetRatingViewAPI />} />*/}
      </Routes>

      <Footer />
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

// Стили для ссылок в навигации
const bottonLinkStyle: React.CSSProperties = {
  color: "#000000",
  textDecoration: "none",
  margin: "0 15px",
  fontSize: "25px",
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
