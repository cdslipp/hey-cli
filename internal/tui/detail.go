package tui

import (
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
)

type detailModel struct {
	viewport viewport.Model
	title    string
	styles   styles
}

func newDetailModel(s styles) detailModel {
	vp := viewport.New(viewport.WithWidth(0), viewport.WithHeight(0))
	return detailModel{viewport: vp, styles: s}
}

func (m *detailModel) setContent(title, body string) {
	m.title = title
	m.viewport.SetContent(body)
	m.viewport.GotoTop()
}

func (m *detailModel) setSize(w, h int) {
	m.viewport.SetWidth(w)
	m.viewport.SetHeight(h - 1) // leave room for title bar
}

func (m detailModel) update(msg tea.Msg) (detailModel, tea.Cmd) {
	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m detailModel) view() string {
	titleBar := m.styles.title.Render(m.title)
	return titleBar + "\n" + m.viewport.View()
}
