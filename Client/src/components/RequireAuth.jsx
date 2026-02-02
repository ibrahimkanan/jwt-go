import { Navigate } from "react-router-dom";
import authStore from "../stores/authStore";

export default function RequireAuth(props) {
  const store = authStore();

  if (store.loggedIn === null) {
    return <div>Loading...</div>;
  }

  if (store.loggedIn === false) {
    return <Navigate to="/login" />;
  }

  return <div>{props.children}</div>;
}
