// +build !ignore_autogenerated

/*
Copyright 2020 The KubeSphere authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// autogenerated by controller-gen object, do not modify manually

package panels

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasePanel) DeepCopyInto(out *BasePanel) {
	*out = *in
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]Target, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasePanel.
func (in *BasePanel) DeepCopy() *BasePanel {
	if in == nil {
		return nil
	}
	out := new(BasePanel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Graph) DeepCopyInto(out *Graph) {
	*out = *in
	in.BasePanel.DeepCopyInto(&out.BasePanel)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Graph.
func (in *Graph) DeepCopy() *Graph {
	if in == nil {
		return nil
	}
	out := new(Graph)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Panel) DeepCopyInto(out *Panel) {
	*out = *in
	in.Graph.DeepCopyInto(&out.Graph)
	in.Table.DeepCopyInto(&out.Table)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Panel.
func (in *Panel) DeepCopy() *Panel {
	if in == nil {
		return nil
	}
	out := new(Panel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Row) DeepCopyInto(out *Row) {
	*out = *in
	if in.Panels != nil {
		in, out := &in.Panels, &out.Panels
		*out = make([]Panel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Row.
func (in *Row) DeepCopy() *Row {
	if in == nil {
		return nil
	}
	out := new(Row)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SingleStat) DeepCopyInto(out *SingleStat) {
	*out = *in
	in.BasePanel.DeepCopyInto(&out.BasePanel)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SingleStat.
func (in *SingleStat) DeepCopy() *SingleStat {
	if in == nil {
		return nil
	}
	out := new(SingleStat)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sort) DeepCopyInto(out *Sort) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sort.
func (in *Sort) DeepCopy() *Sort {
	if in == nil {
		return nil
	}
	out := new(Sort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Style) DeepCopyInto(out *Style) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Style.
func (in *Style) DeepCopy() *Style {
	if in == nil {
		return nil
	}
	out := new(Style)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Table) DeepCopyInto(out *Table) {
	*out = *in
	in.BasePanel.DeepCopyInto(&out.BasePanel)
	out.Sort = in.Sort
	if in.Styles != nil {
		in, out := &in.Styles, &out.Styles
		*out = make([]Style, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Table.
func (in *Table) DeepCopy() *Table {
	if in == nil {
		return nil
	}
	out := new(Table)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Target) DeepCopyInto(out *Target) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Target.
func (in *Target) DeepCopy() *Target {
	if in == nil {
		return nil
	}
	out := new(Target)
	in.DeepCopyInto(out)
	return out
}
