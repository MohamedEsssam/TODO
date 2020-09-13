import React, { useEffect, useState } from "react";
import { useHistory, Link } from "react-router-dom";
import { logout, currentUser } from "../../services/userServices";
import Login from "../../components/login/login";

const Home = () => {
  const history = useHistory();
  const user = currentUser();

  useEffect(() => {});

  return (
    <>
      {!user && (
        <>
          <Link
            to={{
              pathname: "/",
            }}
          >
            login
          </Link>
          ||
          <Link
            to={{
              pathname: "/register",
            }}
          >
            register
          </Link>
          <Login />
        </>
      )}
      {user && (
        <Link
          to={{
            pathname: "/",
          }}
          onClick={logout}
        >
          logout
        </Link>
      )}
    </>
  );
};

export default Home;
