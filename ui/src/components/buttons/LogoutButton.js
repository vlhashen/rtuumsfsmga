import React from "react";
import { instance } from "../../utils/Utils";

const LogoutButton = () => {
  const clickHandler = async () => {
    const res = await instance.get("/logout");
    const status_code = await res.status;
    if (status_code === 200) {
      window.location.reload();
    }
  };

  return (
    <React.Fragment>
      <button
        className="block w-1/2 mx-auto mb-2 font-mono font-bold bg-red-400 border-2 border-black"
        onClick={clickHandler}
      >
        logout
      </button>
    </React.Fragment>
  );
};

export default LogoutButton;
