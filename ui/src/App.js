import React from "react";
import { useState, useEffect } from "react";
import LoginPage from "./components/LoginPage.js";
import FetchButton from "./components/buttons/FetchButton";
import MainSearchBar from "./components/items/MainSearchBar";
import MenuButton from "./components/buttons/MenuButton";
import SavedList from "./components/lists/SavedList";
import MenuModal from "./components/MenuModal";
import SampleList from "./components/lists/SampleList";
import Context from "./store/context";
import { instance } from "./utils/Utils";

const lodash = require("lodash");

export default function App() {
  // FetchedItems data yang diambil
  // SavedItems data yang ditampilkan
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [fetchedItems, setFetchedItems] = useState([]);
  const [savedItems, setSavedItems] = useState([]);
  const [menuVisible, setMenuVisible] = useState(false);
  const [sampleVisible, setSampleVisible] = useState(false);
  const [subredditKeys, setSubredditKeys] = useState([]);
  const [imageDisplayToggle, setImageDisplayToggle] = useState(false);

  useEffect(() => {
    const getData = async () => {
      const result = await instance.get("/all");
      const res_status = await result.status;
      if (res_status === 200 || res_status === 201) {
        setIsAuthenticated(true);
        if (res_status === 201) {
          setFetchedItems(result.data);
          setSavedItems(result.data);
        }
      }
    };
    getData();
  }, []);

  useEffect(() => {
    const slist = lodash.groupBy(fetchedItems, "subreddit");
    const skey = Object.keys(slist).sort();
    setSubredditKeys(skey);
  }, [fetchedItems]);

  useEffect(() => {
    window.scrollTo(0, 0);
  }, [savedItems]);

  return isAuthenticated ? (
    <Context.Provider
      value={{
        fetchedItems,
        setFetchedItems,
        savedItems,
        setSavedItems,
        menuVisible,
        setMenuVisible,
        sampleVisible,
        setSampleVisible,
        subredditKeys,
        isAuthenticated,
        imageDisplayToggle,
        setImageDisplayToggle
      }}
    >
      <div className={menuVisible ? "block" : "hidden"}>
        <MenuModal />
      </div>
      <div className={"float-right p-4 mt-6"}>
        <FetchButton />
        <MenuButton />
      </div>
      <div className="flex font-semibold font-['Helvetica'] tracking-tighter underline decoration-dashed text-6xl place-content-center">
        <h1 className="m-6">rtuumsfsmga</h1>
      </div>

      <MainSearchBar />
      <SampleList />
      <SavedList datas={savedItems} concise={false} />
    </Context.Provider>
  ) : (
    <LoginPage />
  );
}
