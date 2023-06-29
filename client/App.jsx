import "./styles/styles.scss";

import Pokemon from "./components/Pokemon";
import Title from "./components/Title";

export const App = () => {
  return (
    <>
      <div>
        <Title />
        <Pokemon />
      </div>
    </>
  );
};
