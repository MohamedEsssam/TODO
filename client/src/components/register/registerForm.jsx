import React from "react";
import { Form } from "formik";
import { Button } from "react-bootstrap";
import FormInput from "../form/formInput";
import "./registerStyle.css";

const RegisterForm = () => {
  return (
    <>
      <div class="registerContainer">
        <div id="form">
          <h1>Register</h1>
          <Form>
            <div>
              <FormInput
                type="name"
                name="name"
                placeholder="Name"
                id="formInput"
              />
            </div>
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
              Register
            </Button>
          </Form>
        </div>
      </div>
    </>
  );
};

export default RegisterForm;
