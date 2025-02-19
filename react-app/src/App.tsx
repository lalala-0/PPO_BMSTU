import React, { useEffect } from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import { GetRatingsView, RatingView, CrewView, RaceView, ParticipantView } from "./ui";
import StartProcedure from "./ui/view/raceViews/startView";
import FinishProcedure from "./ui/view/raceViews/finishView";
import ProtestView from "./ui/view/protestViews/protestView";
import Login from "./ui/view/autchViews/login";
import { isAuthenticated } from "./ui/controllers/autchControllers/isAutchenticated";
import JudgeProfile from "./ui/view/judgeViews/profile";
import JudgeDashboard from "./ui/view/judgeViews/dashboard";
import OceanAnimation from "./ui/view/ocean";
import Header from "./Header";
import Footer from "./Footer";
import "./index.css";

const App: React.FC = () => {
  useEffect(() => {
    document.body.classList.toggle("authenticated", isAuthenticated());
  }, []);

  return (
      <Router>
        <Header />

        <Routes>
          <Route path="/" element={
            <div>
              <h1>Добро пожаловать!</h1>
              <OceanAnimation />
            </div>}
          />
          <Route path="/participants/:participantID" element={<ParticipantView />} />
          <Route path="/profile" element={<JudgeProfile />} />
          <Route path="/ratings" element={<GetRatingsView />} />
          <Route path="/ratings/:ratingID" element={<RatingView />} />
          <Route path="/ratings/:ratingID/races/:raceID" element={<RaceView />} />
          <Route path="/ratings/:ratingID/races/:raceID/startProcedure" element={<StartProcedure />} />
          <Route path="/ratings/:ratingID/races/:raceID/finishProcedure" element={<FinishProcedure />} />
          <Route path="/ratings/:ratingID/races/:raceID/protests/:protestID" element={<ProtestView />} />
          <Route path="/ratings/:ratingID/crews/:crewID" element={<CrewView />} />
          <Route path="/login" element={<Login />} />
          <Route path="/judgeDashboard" element={<JudgeDashboard />} />
        </Routes>

        <Footer />
      </Router>
  );
};

export default App;
