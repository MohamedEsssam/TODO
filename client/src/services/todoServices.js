import http from "./httpService";
const apiEndpoint = "http://localhost:8000/";

export async function getTodos(userId) {
  const todos = await http.get(apiEndpoint + `${userId}`, {
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
  });

  return todos;
}

export async function getTodo(userId, roomId) {
  const todo = await http.get(apiEndpoint + `${userId}/${roomId}`, {
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
  });

  return todo;
}

export async function create(userId, data) {
  const todo = await http.post(apiEndpoint + `${userId}`, data, {
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
  });

  return todo;
}

export async function update(userId, data) {
  console.log(data);
  const todo = await http.put(apiEndpoint + `${userId}`, data, {
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
  });

  return todo;
}

export async function remove(userId, data) {
  const todo = await http.delete(
    apiEndpoint + `${userId}`,
    { data },
    {
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
    }
  );

  return todo;
}
