import { Box, List, ThemeIcon, Title } from "@mantine/core";
import { CheckCircleFillIcon, XIcon } from "@primer/octicons-react";
import useSWR from "swr";
import "./App.css";
import AddTodo from "./Components/AddTodo";

export interface Todo {
  id: number;
  title: string;
  body: string;
  done: boolean;
}

export const ENDPOINT = "http://localhost:4000";

const fetcher = (url: string) =>
  fetch(`${ENDPOINT}/${url}`).then((r) => r.json());

function App() {
  const { data, mutate } = useSWR<Todo[]>("api/todos", fetcher);

  async function markTodoAdDone(id: number) {
    const updated = await fetch(`${ENDPOINT}/api/todos/${id}/done`, {
      method: "PATCH",
    }).then((r) => r.json());

    mutate(updated);
  }

  async function deleteTodo(id: number) {
    console.log('id: ', id)
    const updated = await fetch(`${ENDPOINT}/api/todos/${id}/delete`, {
      method: "DELETE",
    }).then((r) => r.json());

    mutate(updated);
  }

  return (
    <Box
      sx={(theme) => ({
        padding: "2rem",
        width: "100%",
        maxWidth: "40rem",
        margin: "0 auto",
      })}
    >
      <List spacing="xs" size="sm" mb={12} center >
        {data?.map((todo) => {
          return (
            <List.Item

              key={`todo_list__${todo.id}`}
              sx={(theme) => ({ display: 'flex', justifyContent: 'space-between' })}
              icon={
                todo.done ? (
                  <ThemeIcon color="teal" size={24} radius="xl">
                    <CheckCircleFillIcon size={20} />
                  </ThemeIcon>
                ) : (
                  <ThemeIcon color="gray" size={24} radius="xl"  sx={(theme) => ({ cursor:'pointer' })} onClick={() => markTodoAdDone(todo.id)}>
                    <CheckCircleFillIcon size={20} />
                  </ThemeIcon>
                )
              }
            >
              <Title order={5}>{todo.title}</Title>

              <span onClick={() => deleteTodo(todo.id)}><XIcon /></span>
            </List.Item>
          );
        })}
      </List>

      <AddTodo mutate={mutate} />
    </Box>
  );
}

export default App;