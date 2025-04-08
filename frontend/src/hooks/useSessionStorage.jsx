function useSessionStorage(url) {
  function loadFromSessionStorage() {
    return JSON.parse(window.sessionStorage.getItem(`/${url}`));
  }

  function saveToSessionStorage(data) {
    window.sessionStorage.setItem(`/${url}`, JSON.stringify(data));
  }

  return {
    loadFromSessionStorage,
    saveToSessionStorage,
  };
}

export default useSessionStorage;
