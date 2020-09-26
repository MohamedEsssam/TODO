import React from "react";
import { Formik, Form } from "formik";
import { Button } from "react-bootstrap";
import CustomInput from "./inputComponent";
import todoSchema from "./validation";
import { currentUser } from "../../services/userServices";
import { create } from "../../services/todoServices";

const NewTodo = React.memo(({ initialValues, loadTodos }) => {
  const user = currentUser();
  const onSubmit = async (values, { resetForm }) => {
    const todo = await create(user.userId, values);
    if (todo) {
      resetForm({ values: "" });
      loadTodos();
    }
  };

  return (
    <>
      <Formik
        initialValues={initialValues}
        validationSchema={todoSchema}
        onSubmit={onSubmit}
      >
        {(handleChange, handleSubmit, error, setFieldTouched, touched) => (
          <>
            <Form>
              <CustomInput
                initialValues=""
                type="text"
                name="text"
                placeholder="Enter New Todo"
              />
              <Button type="submit" variant="primary">
                NewTodo
              </Button>
            </Form>
          </>
        )}
      </Formik>
    </>
  );
});

export default NewTodo;
