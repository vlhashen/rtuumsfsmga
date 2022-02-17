import React from "react";
import { instance } from "../utils/Utils";

const LoginPage = () => {
  const loginHandler = async () => {
    const result = await instance.get("/login");
    const auth_url = result.data.auth_url;
    window.location.replace(auth_url);
    //window.history.pushState("", "", auth_url);
    //window.location.reload();
  };

  return (
    <React.Fragment>
      <div className="flex font-semibold font-['Helvetica'] tracking-tighter underline decoration-dashed text-6xl place-content-center">
        <h1 className="mt-6 mb-40">rtuumsfsmga</h1>
      </div>
      <div className="w-1/3 ml-auto mr-auto ">
        <button
          onClick={loginHandler}
          className="items-center ml-auto mr-auto border-4 border-black border-dashed hover:border-solid grid grid-cols-2 gap-2"
        >
          <p className="font-mono font-bold underline ">login</p>
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="40"
            height="40"
            viewBox="0 0 24 24"
            className="ml-auto mr-0 fill-red-400"
          >
            <path d="M14.558 15.827c.097.096.097.253 0 .349-.531.529-1.365.786-2.549.786l-.009-.002-.009.002c-1.185 0-2.018-.257-2.549-.786-.097-.096-.097-.253 0-.349.096-.096.254-.096.351 0 .433.431 1.152.641 2.199.641l.009.002.009-.002c1.046 0 1.765-.21 2.199-.641.095-.097.252-.097.349 0zm-.126-3.814c-.581 0-1.054.471-1.054 1.05 0 .579.473 1.049 1.054 1.049.581 0 1.054-.471 1.054-1.049 0-.579-.473-1.05-1.054-1.05zm9.568-12.013v24h-24v-24h24zm-4 11.853c0-.972-.795-1.764-1.772-1.764-.477 0-.908.191-1.227.497-1.207-.794-2.84-1.299-4.647-1.364l.989-3.113 2.677.628-.004.039c0 .795.65 1.442 1.449 1.442.798 0 1.448-.647 1.448-1.442 0-.795-.65-1.442-1.448-1.442-.613 0-1.136.383-1.347.919l-2.886-.676c-.126-.031-.254.042-.293.166l-1.103 3.471c-1.892.023-3.606.532-4.867 1.35-.316-.292-.736-.474-1.2-.474-.975-.001-1.769.79-1.769 1.763 0 .647.355 1.207.878 1.514-.034.188-.057.378-.057.572 0 2.607 3.206 4.728 7.146 4.728 3.941 0 7.146-2.121 7.146-4.728 0-.183-.019-.362-.05-.54.555-.299.937-.876.937-1.546zm-9.374 1.21c0-.579-.473-1.05-1.054-1.05-.581 0-1.055.471-1.055 1.05 0 .579.473 1.049 1.055 1.049.581.001 1.054-.47 1.054-1.049z" />
          </svg>
        </button>
      </div>
    </React.Fragment>
  );
};

export default LoginPage;
