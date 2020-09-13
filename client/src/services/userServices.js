import http from "./httpService";
const apiEndpoint = "http://localhost:8000/";

export async function login(data) {
  const user = await http.post(apiEndpoint + "login", data, {
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
  });

  localStorage.setItem("user", JSON.stringify(user.data));

  return currentUser();
}

export async function register(data) {
  const user = await http.post(apiEndpoint + "register", data, {
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
  });

  localStorage.setItem("user", JSON.stringify(user.data));

  return currentUser();
}

export function logout() {
  localStorage.removeItem("user");
}

export function currentUser() {
  return JSON.parse(localStorage.getItem("user"));
}
