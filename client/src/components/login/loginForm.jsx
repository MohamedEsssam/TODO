import React from "react";
import { Button } from "react-bootstrap";
import { Form, useFormikContext } from "formik";
import FormInput from "../form/formInput";

const LoginForm = () => {
  const { handleChange, setFieldTouched, touched, errors } = useFormikContext();
  return (
    <div id="form">
      <h1>Login</h1>
      <Form>
        <div>
          <FormInput
            type="email"
            name="email"
            placeholder="Email"
            id="formInput"
          />
        </div>
        <div>
          <FormInput
            type="password"
            name="password"
            placeholder="Password"
            id="formInput"
          />
        </div>
        <Button
          type="submit"
          variant="dark"
          size="lg"
          className="mr-2"
          id="login"
        >
          Login
        </Button>
      </Form>
    </div>
  );
};

export default LoginForm;
