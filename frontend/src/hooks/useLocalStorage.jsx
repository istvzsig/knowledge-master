function useLocalStorage(url) {
  function loadFromLocalStore() {
    return JSON.parse(window.localStorage.getItem(`/${url}`));
  }

  function saveTolocalStorage(data) {
    window.localStorage.setItem(`/${url}`, JSON.stringify(data));
  }

  return {
    loadFromLocalStore,
    saveTolocalStorage,
  };
}

export default useLocalStorage;
