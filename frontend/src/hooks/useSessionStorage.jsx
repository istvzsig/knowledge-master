function useSessionStorage() {
  function load(path) {
    return JSON.parse(window.sessionStorage.getItem(path));
  }

  function save(path, data) {
    window.sessionStorage.setItem(path, JSON.stringify(data));
  }

  return {
    load,
    save,
  };
}

export default useSessionStorage;
