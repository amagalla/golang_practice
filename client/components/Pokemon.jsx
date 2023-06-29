import "../styles/pokemon.scss";
import { useState, useEffect } from "react";
import axios from "axios";

const Pokemon = () => {
  const [pokemonList, setPokemonList] = useState([]);
  const [isVisible, setIsVisible] = useState(false);
  const [pageConfig, setPageConfig] = useState({
    page: 1,
    page_size: 3,
    totalPages: 0,
  });

  const getPokemonApi = async () => {
    try {
      const resp = await axios.get(
        `http://localhost:3000/pokemon?page=${pageConfig.page}&page_size=${pageConfig.page_size}`
      );
      return resp;
    } catch (err) {
      return err;
    }
  };

  const pokemonHandleClick = async () => {
    toggleVisibility();
    const resp = await getPokemonApi();

    setPokemonList(resp.data.pokemonList);
    setPageConfig((prevState) => ({
      ...prevState,
      totalPages: resp.data.totalPages,
    }));
  };

  const leftHandleClick = async () => {
    setPageConfig((prevState) => ({
      ...prevState,
      page: prevState.page - 1,
    }));
  };

  const rightHandleClick = async () => {
    setPageConfig((prevState) => ({
      ...prevState,
      page: prevState.page + 1,
    }));
  };

  useEffect(() => {
    const fetchData = async () => {
      const resp = await getPokemonApi();
      setPokemonList(resp.data.pokemonList);
    };

    if (
      pageConfig.page >= 1 &&
      pageConfig.page <= pageConfig.totalPages
    ) {
      fetchData();
    } else {
      setPageConfig((prevState) => ({
        ...prevState,
        page: Math.max(1, Math.min(pageConfig.page, pageConfig.totalPages)),
      }));
    }
  }, [pageConfig.page, pageConfig.totalPages]);

  const toggleVisibility = () => {
    setIsVisible(!isVisible);
  };

  const pokemonListOutput = pokemonList.map((pokemon, i) => (
    <div className="pokemon-names" key={i}>
      {pokemon.name}
    </div>
  ));

  return (
    <>
      <div className="pokemon-button">
        <button onClick={pokemonHandleClick}>Click for Pokemon</button>
      </div>
      {isVisible && (
        <>
          <div className="pokemon-container">{pokemonListOutput}</div>
          <div className="pokemon-pagination">
            <button
              onClick={leftHandleClick}
              disabled={pageConfig.page <= 1}
            >
              Left
            </button>
            <div>{pageConfig.page} of {pageConfig.totalPages}</div>
            <button
              onClick={rightHandleClick}
              disabled={pageConfig.page >= pageConfig.totalPages}
            >
              Right
            </button>
          </div>
        </>
      )}
    </>
  );
};

export default Pokemon;
