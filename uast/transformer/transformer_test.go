package transformer

import (
	"testing"

	"github.com/stretchr/testify/require"
	u "gopkg.in/bblfsh/sdk.v2/uast"
	un "gopkg.in/bblfsh/sdk.v2/uast/nodes"
)

var mappingCases = []struct {
	name     string
	skip     bool
	inp, exp un.Node
	m        Transformer
}{
	{
		name: "trim meta",
		inp: un.Object{
			"the_root": un.Object{
				"k": un.String("v"),
			},
		},
		m: ResponseMetadata{
			TopLevelIsRootNode: false,
		},
		exp: un.Object{
			"k": un.String("v"),
		},
	},
	{
		name: "leave meta",
		inp: un.Object{
			"the_root": un.Object{
				"k": un.String("v"),
			},
		},
		m: ResponseMetadata{
			TopLevelIsRootNode: true,
		},
	},
	{
		name: "roles dedup",
		inp: un.Array{
			un.Object{
				u.KeyType:  un.String("typed"),
				u.KeyRoles: u.RoleList(1, 2, 1),
			},
		},
		m: RolesDedup(),
		exp: un.Array{
			un.Object{
				u.KeyType:  un.String("typed"),
				u.KeyRoles: u.RoleList(1, 2),
			},
		},
	},
	{
		name: "typed and generic",
		inp: un.Array{
			un.Object{
				u.KeyType: un.String("typed"),
				"pred":    un.String("val1"),
				"k":       un.String("v"),
			},
			un.Object{
				"pred": un.String("val2"),
				"k":    un.String("v"),
			},
			un.Object{
				"pred2": un.String("val3"),
			},
		},
		m: Mappings(
			Map(
				Part("_", Obj{
					"pred": Var("x"),
				}),
				Part("_", Obj{
					"p": Var("x"),
				}),
			),
			AnnotateType("typed", MapObj(Obj{
				"k": Var("x"),
			}, Obj{
				"key": Var("x"),
			}), 10),
		),
		exp: un.Array{
			un.Object{
				u.KeyType:  un.String("typed"),
				u.KeyRoles: u.RoleList(10),
				"p":        un.String("val1"),
				"key":      un.String("v"),
			},
			un.Object{
				"p": un.String("val2"),
				"k": un.String("v"),
			},
			un.Object{
				"pred2": un.String("val3"),
			},
		},
	},
	{
		name: "annotate no roles",
		inp: un.Array{
			un.Object{
				u.KeyType:  un.String("typed"),
				u.KeyRoles: u.RoleList(1),
				"pred":     un.String("val1"),
			},
			un.Object{
				u.KeyType: un.String("typed"),
				"pred":    un.String("val2"),
			},
		},
		m: Mappings(AnnotateIfNoRoles("typed", 10)),
		exp: un.Array{
			un.Object{
				u.KeyType:  un.String("typed"),
				u.KeyRoles: u.RoleList(1),
				"pred":     un.String("val1"),
			},
			un.Object{
				u.KeyType:  un.String("typed"),
				u.KeyRoles: u.RoleList(10),
				"pred":     un.String("val2"),
			},
		},
	},
}

func TestMappings(t *testing.T) {
	for _, c := range mappingCases {
		if c.exp == nil {
			c.exp = c.inp
		}
		t.Run(c.name, func(t *testing.T) {
			out, err := c.m.Do(c.inp)
			require.NoError(t, err)
			require.Equal(t, c.exp, out, "transformation failed")
		})
	}
}
