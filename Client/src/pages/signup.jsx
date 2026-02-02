import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { toast } from "react-hot-toast";
import authStore from "../stores/authStore";
import "../styles/Auth.css";

function Signup() {
  const store = authStore();
  const navigate = useNavigate();
  const [formData, setFormData] = useState({ username: "", password: "" });

  const handleSignup = async (e) => {
    e.preventDefault();
    try {
      await store.signup(formData);
      toast.success("Signup successful!");
      navigate("/login");
    } catch (error) {
      console.log(error);
      const errorMessage =
        error.response?.data?.error || error.message || "Signup failed";
      toast.error(errorMessage);
    }
  };
  return (
    <div>
      <div className="auth-container">
        <div className="auth-header">
          <h1>Create Account</h1>
          <p>Join us today and start your journey</p>
        </div>

        <form className="auth-form" onSubmit={handleSignup}>
          <div className="form-group">
            <input
              className="form-input"
              type="text"
              placeholder="Username"
              value={formData.username}
              onChange={(e) =>
                setFormData({ ...formData, username: e.target.value })
              }
              required
            />
          </div>

          <div className="form-group">
            <input
              className="form-input"
              type="password"
              placeholder="Password"
              value={formData.password}
              onChange={(e) =>
                setFormData({ ...formData, password: e.target.value })
              }
              required
            />
          </div>

          <button className="submit-btn" type="submit">
            Sign Up
          </button>
        </form>

        <div className="auth-footer">
          <p>
            Already have an account? <Link to="/login">Login</Link>
          </p>
        </div>
      </div>
    </div>
  );
}

export default Signup;
