window.onload = function () {
  createTable(13, 13)

  // Number, Row, Cell, Count, toDown
  // #1
  insertCell(1, 0, 0, 6, true)
  insertCell(1, 0, 0, 6, false)
  // #2
  insertCell(2, 0, 3, 6, true)
  // #3
  insertCell(3, 0, 5, 4, true)
  // #4
  insertCell(4, 0, 8, 7, true)
  // #5
  insertCell(5, 0, 10, 5, true)
  // #6
  insertCell(6, 0, 12, 5, true)
  // #7
  insertCell(7, 1, 5, 8, false)
  // #8
  insertCell(8, 2, 2, 4, true)
  insertCell(8, 2, 2, 4, false)
  // #9
  insertCell(9, 2, 4, 3, true)
  // #10
  insertCell(10, 3, 0, 6, false)
  // #11
  insertCell(11, 3, 7, 6, false)
  // #12
  insertCell(12, 3, 9, 3, true)
  // #13
  insertCell(13, 3, 11, 5, true)
  // #14
  insertCell(14, 4, 2, 3, false)
  // #15
  insertCell(15, 4, 6, 5, true)
  // #16
  insertCell(16, 4, 8, 5, false)
  // #17
  insertCell(17, 5, 0, 4, false)
  // #18
  insertCell(18, 5, 1, 5, true)
  // #19
  insertCell(19, 5, 5, 5, false)
  // #20
  insertCell(20, 5, 7, 3, true)
  // #21
  insertCell(21, 6, 4, 7, true)
  insertCell(21, 6, 4, 5, false)
  // #22
  insertCell(22, 7, 3, 3, true)
  insertCell(22, 7, 3, 5, false)
  // #23
  insertCell(23, 7, 9, 6, true)
  insertCell(23, 7, 9, 4, false)
  // #24
  insertCell(24, 7, 10, 4, true)
  // #25
  insertCell(25, 7, 12, 6, true)
  // #26
  insertCell(26, 8, 0, 5, true)
  insertCell(26, 8, 0, 5, false)
  // #27
  insertCell(27, 8, 2, 5, true)
  // #28
  insertCell(28, 8, 8, 3, true)
  insertCell(28, 8, 8, 3, false)
  // #29
  insertCell(29, 9, 0, 6, false)
  // #30
  insertCell(30, 9, 7, 6, false)
  insertCell(30, 9, 7, 4, true)
  // #31
  insertCell(31, 10, 7, 4, false)
  // #32
  insertCell(32, 11, 0, 8, false)
  // #33
  insertCell(32, 12, 7, 6, false)
}

function createTable(rows, cells) {
  // Declare Table ID
  let table = document.getElementById('tableTts')

  // Create Table Row
  for (let row = 0; row < rows; row++) {
    // Declare Table Row
    let tableRow = table.insertRow(row)
    // Create Table Cell
    for (let cell = 0; cell < cells; cell++) {
      tableRow.insertCell(cell)
    }
  }
}

function insertCell(number, row, cell, count, toDown) {
  // Declare Table ID
  let table = document.getElementById('tableTts')

  for (let i = 0; i < count; i++) {
    if (toDown) {
      table.rows[row + i].cells[cell].style.backgroundColor = '#ffffff'
      table.rows[row + i].cells[cell].innerHTML = '<input type="text" class="tts-text" maxlength="1">'
    } else {
      table.rows[row].cells[cell + i].style.backgroundColor = '#ffffff'
      table.rows[row].cells[cell + i].innerHTML = '<input type="text" class="tts-text" maxlength="1">'
    }
  }

  table.rows[row].cells[cell].innerHTML += '<text>' + number + '</text>'
}
