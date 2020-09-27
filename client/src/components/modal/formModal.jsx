import React from "react";
import { Formik, Form } from "formik";
import { toast } from "react-toastify";
import { Button } from "react-bootstrap";
import TodoSchema from "./validation";
import { update } from "../../services/todoServices";
import { currentUser } from "../../services/userServices";
import FormInput from "../form/formInput";

const FromModal = React.memo(({ initialValues, todoId, onHide }) => {
  const onSubmit = async (values) => {
    const user = currentUser();
    values.todoId = todoId;
    const todo = await update(user.userId, values);

    if (todo) onHide();
  };

  return (
    <>
      <Formik
        initialValues={initialValues}
        validationSchema={TodoSchema}
        onSubmit={onSubmit}
      >
        {(handleChange, handleSubmit, error, setFieldTouched, touched) => (
          <>
            <Form>
              <FormInput
                initialValues={initialValues}
                type="text"
                name="text"
                placeholder="Update Todo"
              />
              <Button type="submit" variant="primary">
                Update Todo
              </Button>
            </Form>
          </>
        )}
      </Formik>
    </>
  );
});

export default FromModal;
