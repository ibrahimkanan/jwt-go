import React from "react";
import authStore from "../stores/authStore";
import { useNavigate } from "react-router-dom";
import "../styles/Mainpage.css";

export default function Mainpage() {
  const store = authStore();
  const navigate = useNavigate();

  const handleLogout = async () => {
    await store.logout();
    navigate("/login");
  };

  return (
    <div className="main-container">
      <header className="dashboard-header">
        <h1>Welcome to Dashboard</h1>
        <button className="logout-btn" onClick={handleLogout}>
          Logout
        </button>
      </header>

      <div className="content-grid">
        <div className="card">
          <h3>Your Profile</h3>
          <p>Manage your personal information and settings.</p>
        </div>

        <div className="card">
          <h3>Analytics</h3>
          <p>View your activity and performance metrics.</p>
        </div>

        <div className="card">
          <h3>Messages</h3>
          <p>Check your latest notifications and messages.</p>
        </div>

        <div className="card">
          <h3>Settings</h3>
          <p>Configure your application preferences.</p>
        </div>
      </div>
    </div>
  );
}
