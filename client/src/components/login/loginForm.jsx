import React from "react";
import { Button } from "react-bootstrap";
import { Form, useFormikContext } from "formik";
import CustomInput from "../forms/inputComponent";

const LoginForm = () => {
  const { handleChange, setFieldTouched, touched, errors } = useFormikContext();
  return (
    <div id="form">
      <h1>Login</h1>
      <Form>
        <div>
          <CustomInput
            type="email"
            name="email"
            placeholder="Email"
            id="formInput"
          />
        </div>
        <div>
          <CustomInput
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
