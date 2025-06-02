import XCTest
import SwiftTreeSitter
import TreeSitterOneScript

final class TreeSitterOneScriptTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_onescript())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading OneScript grammar")
    }
}
