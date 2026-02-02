import { Routes, Route } from "react-router-dom";
import { Toaster } from "react-hot-toast";
import { useEffect } from "react";
import Home from "./pages/Home";
import Login from "./pages/Login";
import Signup from "./pages/signup";
import Mainpage from "./pages/Mainpage";
import RequireAuth from "./components/RequireAuth";
import authStore from "./stores/authStore";
import "./App.css";

function App() {
  const checkAuth = authStore((state) => state.checkAuth);

  useEffect(() => {
    checkAuth();
  }, [checkAuth]);

  return (
    <div className="App">
      <Toaster position="top-center" />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/signup" element={<Signup />} />
        <Route
          path="/main"
          element={
            <RequireAuth>
              <Mainpage />
            </RequireAuth>
          }
        />
      </Routes>
    </div>
  );
}

export default App;
