generateTitle = () => {
  return 'Title';
}

createTag = (title = generateTitle(), tag = 'css') => {
  console.log('title: ' + title + ' tag: ' + tag);
}

createTag();