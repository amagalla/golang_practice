import "../styles/pokemon.scss";
import { useState } from "react";
import axios from "axios";

const Pokemon = () => {
  const [pokemonList, setPokemonList] = useState([]);
  const [isVisible, setIsVisible] = useState(false);

  const pokemonHandleClick = async () => {
    toggleVisibility();
    try {
      const response = await axios.get("http://localhost:3000/pokemon");

      setPokemonList(response.data.pokemonList);
    } catch (err) {
      console.log("err: ", err.response);
    }
  };

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
        <div className="pokemon-container">{pokemonListOutput}</div>
      )}
    </>
  );
};

export default Pokemon;
